"use client";

import { ComponentPropsWithRef, RefObject, useImperativeHandle, useRef } from "react";

export interface InterviewChatSoundControllerRef {
  playNewMessageSound: () => void;
  playInterviewCompletedSound: () => void;
}

export function InterviewChatSoundController({ ref }: ComponentPropsWithRef<any>) {
  const newMessageSound = useRef<HTMLAudioElement>(null);
  const interviewCompletedSound = useRef<HTMLAudioElement>(null);

  const playSoundHandler = (soundRef: RefObject<HTMLAudioElement | null>) => {
    if (soundRef.current !== null) {
      soundRef.current.currentTime = 0;
      soundRef.current.play().catch((error) => {
        console.error("Error playing sound:", error);
      });
    }
  };

  useImperativeHandle(ref, () => ({
    playNewMessageSound: () => playSoundHandler(newMessageSound),
    playInterviewCompletedSound: () => playSoundHandler(interviewCompletedSound),
  }));

  return (
    <>
      <audio ref={newMessageSound} preload="auto" className="hidden">
        <source src="/sound/new-message-sound.mp3" type="audio/mpeg" />
        Your browser does not support the audio element.
      </audio>

      <audio ref={interviewCompletedSound} preload="auto" className="hidden">
        <source src="/sound/interview-completed-sound.mp3" type="audio/mpeg" />
        Your browser does not support the audio element.
      </audio>
    </>
  );
}
