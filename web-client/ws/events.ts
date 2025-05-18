export enum EventType {
  ClientMessageSent = "client_message_sent",

  SystemMessageSent = "system_message_sent",
  SystemMessagePending = "system_message_pending",
}

export interface BaseEvent {
  type: EventType;
}

export interface ClientMessageSentEvent extends BaseEvent {
  type: EventType.ClientMessageSent;
  content: string;
}

export interface SystemMessageSentEvent extends BaseEvent {
  type: EventType.SystemMessageSent;
  content_text: string;
  tips_text: string;
}

export interface SystemMessagePendingEvent extends BaseEvent {
  type: EventType.SystemMessagePending;
}
