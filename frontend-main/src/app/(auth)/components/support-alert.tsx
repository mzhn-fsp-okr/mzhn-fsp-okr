import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { AlertDialogTriggerProps } from "@radix-ui/react-alert-dialog";

export default function SupportAlert(props: AlertDialogTriggerProps) {
  return (
    <AlertDialog>
      <AlertDialogTrigger {...props} />
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Упс!</AlertDialogTitle>
        </AlertDialogHeader>
        <AlertDialogDescription>
          Пожалуйста, обратитесь к техническому отделу для того, чтобы сделать
          это.
        </AlertDialogDescription>
        <AlertDialogFooter>
          <AlertDialogCancel>Закрыть</AlertDialogCancel>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
