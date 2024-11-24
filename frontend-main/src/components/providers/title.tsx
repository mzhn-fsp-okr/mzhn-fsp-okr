"use client";

import {
  createContext,
  HTMLAttributes,
  PropsWithChildren,
  useContext,
  useEffect,
  useState,
} from "react";

const TitleContext = createContext<{
  title: string;
  setTitle: (value: string) => void;
}>({
  title: "",
  setTitle: () => {},
});

export const TitleProvider = ({ children }: PropsWithChildren) => {
  const [title, setTitle] = useState("");
  return (
    <TitleContext.Provider value={{ title, setTitle }}>
      {children}
    </TitleContext.Provider>
  );
};

export function useTitleContext() {
  return useContext(TitleContext);
}

export function Title(props: HTMLAttributes<HTMLDivElement>) {
  const { title } = useTitleContext();
  return <div {...props}>{title}</div>;
}

export function TitleInit({ title }: { title: string }) {
  const { setTitle } = useTitleContext();
  // eslint-disable-next-line react-hooks/exhaustive-deps -- It will never be updated cause it stored in context
  useEffect(() => setTitle(title), []);
  return <></>;
}
