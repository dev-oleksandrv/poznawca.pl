export interface APIResponseWrapper<T> {
  data: T;
  error?: string;
}
