import { InterviewMessageDto } from "@/dto/interview-dto";

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
  details: InterviewMessageDto;
}

export interface SystemMessagePendingEvent extends BaseEvent {
  type: EventType.SystemMessagePending;
}
