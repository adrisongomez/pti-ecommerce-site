import { FC, useState } from "react";
import { MessageCircle, Star, Terminal } from "react-feather";
import { joinClass } from "../globals/utilities/joinClass";
import Textfield from "../globals/components/fields/Textfield";
import Button from "../globals/components/buttons/Button";
import { useChatServicePostApiChatsBySessionIdMessage } from "../globals/generated/queries";
import { ReactNode } from "@tanstack/react-router";
import LoadingIndicator from "../globals/components/progress/LoadingIndicator";

const WELCOM_MESSAGE = () => (
  <p>
    Welcome to <b>AI Store Assistant</b>! We're here to help you find what you
    need quickly and easily. Whether you're browsing, comparing, or checking
    out, we've got your back. Let’s make shopping smarter! 🛍️✨
  </p>
);

type MessageSource = "user" | "assistant";
const AssitantChat: FC<{
  onOpen: () => Promise<void>;
  onClose: () => Promise<void>;
  sessionId: number;
  open: boolean;
}> = ({ open, sessionId, onOpen, onClose }) => {
  const [prompt, setPrompt] = useState("");
  const [messages, setMessages] = useState<
    { message: string; source: MessageSource }[]
  >([]);

  const sendPrompt = useChatServicePostApiChatsBySessionIdMessage({
    mutationKey: ["AssitantChat_SendPrompt"],
  });

  async function submitMessage() {
    setMessages((messages) => [
      ...messages,
      { message: prompt, source: "user" },
    ]);
    const newMessage = await sendPrompt.mutateAsync({
      sessionId,
      requestBody: { message: prompt },
    });
    setMessages((messages) => [
      ...messages,
      { message: newMessage, source: "assistant" },
    ]);
  }
  return (
    <>
      <button
        onClick={onOpen}
        data-modal-target="default-modal"
        data-modal-toggle="default-modal"
        className={joinClass(
          "fixed right-4 bottom-4 rounded-full bg-(--bg-dark) px-4 py-4 font-bold text-white",
          "shadow-lg transition-all duration-300 ease-in-out hover:bg-(--bg-light) hover:text-(--bg-dark)",
          "outline-none",
        )}
        type="button"
      >
        <MessageCircle />
      </button>

      <div
        id="default-modal"
        tabIndex={1}
        onClick={onClose}
        aria-hidden="true"
        className={joinClass(
          open ? "flex" : "hidden",
          "fixed top-0 right-0 left-0 z-50 h-[calc(100%-1rem)]",
          "bg-black/50 backdrop-blur-md",
          "h-screen w-screen items-center justify-center overflow-x-hidden overflow-y-auto md:inset-0",
        )}
      >
        <div
          onClick={(e) => {
            e.stopPropagation();
          }}
          className="relative max-h-full w-full max-w-2xl p-4"
        >
          <div className="relative rounded-md bg-(--bg-main) shadow-sm">
            <div className="flex items-center justify-between rounded-t border-b border-gray-200 p-4 md:p-5 dark:border-gray-600">
              <h3 className="flex items-center gap-3 text-xl font-semibold text-gray-900">
                <Star size="20" color="var(--bg-dark)" /> AI Store Assistant
              </h3>
              <button
                type="button"
                className="ms-auto inline-flex h-8 w-8 items-center justify-center rounded-lg bg-transparent text-sm text-gray-400 hover:bg-gray-200 hover:text-gray-900 dark:hover:bg-gray-600 dark:hover:text-white"
                data-modal-hide="default-modal"
              >
                <svg
                  onClick={onClose}
                  className="h-3 w-3"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 14 14"
                >
                  <path
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                  />
                </svg>
                <span onClick={onClose} className="sr-only">
                  Close modal
                </span>
              </button>
            </div>
            <div className="flex h-[65vh] flex-col space-y-4 overflow-y-auto p-4 md:p-5">
              {messages.length === 0 && (
                <MessageBox message={<WELCOM_MESSAGE />} source="assistant" />
              )}
              {messages.map((v, i) => (
                <MessageBox
                  key={`message-box-${i}`}
                  message={v.message}
                  source={v.source}
                />
              ))}
              {sendPrompt.isPending && (
                <div className="mx-auto flex-1">
                  <LoadingIndicator />
                </div>
              )}
            </div>
            <div className="flex items-end gap-6 rounded-b p-4 md:p-5">
              <Textfield
                multiple
                fullWidth
                label="Prompt"
                onChange={(e) => {
                  setPrompt(e.target.value ?? "");
                }}
              />{" "}
              <Button
                variant="contained"
                className="rounded-md"
                onClick={submitMessage}
              >
                <Terminal />
              </Button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

const MessageBox: FC<{ message: ReactNode; source: MessageSource }> = ({
  message,
  source,
}) => (
  <div
    className={joinClass(
      source === "assistant" ? "ml-auto" : "mr-auto",
      "mt-6 max-w-md space-y-2",
    )}
  >
    <div
      className={joinClass(
        source === "assistant"
          ? "bg-(--bg-dark) text-white"
          : "bg-(--bg-light) text-black",
        "w-full rounded-lg p-3 shadow-md",
      )}
    >
      {message}
    </div>
  </div>
);

export default AssitantChat;
