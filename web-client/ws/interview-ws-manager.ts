import { BaseInterviewEvent, InterviewEventType } from "@/ws/interview-events";

export class InterviewWsManager {
  private readonly socketConn: WebSocket;
  private readonly listeners: Map<InterviewEventType, Set<Function>> = new Map();

  constructor(private readonly interviewId: string) {
    this.socketConn = new WebSocket(
      `${process.env.NEXT_PUBLIC_INTERVIEW_WS_URL}?interview_id=${interviewId}`,
    );

    this.socketConn.addEventListener("open", () => {
      this.log("WebSocket connection opened");
    });

    this.socketConn.addEventListener("message", (event) => {
      this.handleServerMessage(event);
    });

    this.socketConn.addEventListener("close", () => {
      this.log("WebSocket connection closed");
    });
  }

  public sendEvent<T extends BaseInterviewEvent>(event: T): this {
    try {
      const eventData = JSON.stringify(event);
      this.socketConn.send(eventData);
    } catch (error) {
      this.logError(new Error(`Failed to send event: ${event.type}, Error: ${error}`));
    }
    return this;
  }

  public subscribe(eventName: InterviewEventType, listener: Function): this {
    const existingListeners = this.listeners.get(eventName);
    if (!existingListeners) {
      this.listeners.set(eventName, new Set([listener]));
    } else {
      existingListeners.add(listener);
      this.listeners.set(eventName, existingListeners);
    }

    return this;
  }

  private handleServerMessage(event: MessageEvent) {
    if (!event.data) {
      return;
    }

    try {
      const parsedEvent: BaseInterviewEvent = JSON.parse(event.data);
      if (parsedEvent && parsedEvent.type) {
        const listeners = this.listeners.get(parsedEvent.type);
        if (listeners) {
          listeners.forEach((listener) => listener(parsedEvent));
        } else {
          this.log(`No listeners for event type: ${parsedEvent.type}`);
        }
      }
    } catch (err) {
      this.logError(new Error(`Failed to parse event data: ${event.data}`));
      return;
    }
  }

  private log(message: string, ...args: any[]) {
    console.log(`[InterviewWsManager] ${message}`, ...args);
  }

  private logError(error: Error) {
    console.error(`[InterviewWsManager] Error: ${error.message}`, error);
  }
}
