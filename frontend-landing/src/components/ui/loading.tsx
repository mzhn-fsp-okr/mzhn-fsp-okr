import { LoaderCircle } from "lucide-react";
import { forwardRef } from "react";
import { Button, ButtonProps } from "./button";

export interface LoadingButtonProps extends ButtonProps {
  loading?: boolean;
}

const LoadingButton = forwardRef<HTMLButtonElement, LoadingButtonProps>(
  ({ loading, children, ...props }, ref) => {
    return (
      <Button ref={ref} disabled={loading} {...props}>
        {loading && <LoaderCircle className="size-16 animate-spin" />}
        {!loading && children}
      </Button>
    );
  }
);
LoadingButton.displayName = "LoadingButton";

export { LoadingButton };
