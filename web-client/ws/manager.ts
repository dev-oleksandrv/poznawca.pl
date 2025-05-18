import { InterviewDto } from "@/dto/interview-dto";
import { BaseEvent, ClientMessageSentEvent, EventType, SystemMessageSentEvent } from "@/ws/events";

export enum WSManagerEventType {
  SystemMessageSent = "system_message_sent",
}

export class WSManager {
  private socket: WebSocket;
  private readonly listeners: Map<WSManagerEventType, Function> = new Map();

  constructor(private readonly interview: InterviewDto) {
    this.socket = new WebSocket(`ws://localhost:8080/ws/interview?interview_id=${interview.id}`);

    this.socket.addEventListener("open", () => {
      console.log("[WSManager] WebSocket connection opened");
    });

    this.socket.addEventListener("message", (event) => {
      console.log("[WSManager] Message received:", event.data);
      this.handleMessage(event);
    });

    this.socket.addEventListener("close", () => {
      console.log("[WSManager] WebSocket connection closed");
    });
  }

  public on(event: WSManagerEventType, callback: Function) {
    this.listeners.set(event, callback);
  }

  public sendMessage(message: string) {
    const event: ClientMessageSentEvent = {
      type: EventType.ClientMessageSent,
      content: message,
    };
    this.socket.send(JSON.stringify(event));
  }

  public close(): void {
    this.socket.close();
  }

  private handleMessage(event: MessageEvent) {
    let parsedEvent: BaseEvent;
    try {
      parsedEvent = JSON.parse(event.data);
    } catch (error) {
      console.error("[WSManager] Error parsing message:", error);
      return;
    }

    switch (parsedEvent.type) {
      case EventType.SystemMessageSent:
        this.handleSystemMessageSent(parsedEvent as SystemMessageSentEvent);
        break;
    }
  }

  private handleSystemMessageSent(event: SystemMessageSentEvent) {
    const listener = this.listeners.get(WSManagerEventType.SystemMessageSent);
    if (listener) {
      listener(event);
    } else {
      console.warn("[WSManager] No listener for SystemMessageSent event");
    }
  }
}
