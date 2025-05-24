import { InterviewResultDto } from "@/dto/interview-result-dto";

export const mockResult: InterviewResultDto = {
  id: "a8da0907-4e51-4c65-aa21-e0fbd16b5de9",
  grammar_score: 55,
  accuracy_score: 50,
  total_score: 53,
  grammar_feedback:
    "W Pana odpowiedziach były widoczne drobne błędy językowe, takie jak nieprawidłowe formy słów: 'polakem' zamiast 'Polakiem', 'dobrzym' zamiast 'dobrym', 'Polsce' zamiast 'Polskę'. Proszę zwrócić uwagę na poprawną odmianę słów i składnię zdań.",
  accuracy_feedback:
    "Odpowiedzi były dość ogólne i brakowało im szczegółowości. Warto bardziej rozwinąć odpowiedzi, aby lepiej ukazać swoje zainteresowania i plany związane z Polską.",
  total_feedback:
    "Proszę poświęcić więcej uwagi na naukę gramatyki oraz rozwinięcie odpowiedzi, aby były bardziej szczegółowe i precyzyjne. Ważne jest, aby pokazać swoje zaangażowanie i zrozumienie polskiej kultury i społeczeństwa.",
};
