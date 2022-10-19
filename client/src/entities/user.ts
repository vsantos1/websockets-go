export type User = {
  id: string;
  name: string;
  email: string;
  password: string;
  state: State;
};
type State = {
  status: "online" | "offline" | "busy";
  typing: boolean;
};
