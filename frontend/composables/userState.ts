import { User } from "~/types/user";
import type { Ref } from "vue";
export const useLoginUser = () => {
  const user = useState<User>("login-user", () => {
    return {
      id: 0,
      name: "",
    };
  });

  return {
    user: readonly(user),
    setUser: setUser(user),
  };
};

export const setUser = (user: Ref<User>) => (response: Ref<User>) => {
  user.value.id = response.value.id;
  user.value.name = response.value.name;
};
