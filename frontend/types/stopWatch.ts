import { Ref } from "vue";
export type StopWatch = {
  time: Ref<number>;
  status: Ref<number>;
  startTime: Ref<number>;
  stopTime: Ref<number>;
  timer: Ref<number>;
};
