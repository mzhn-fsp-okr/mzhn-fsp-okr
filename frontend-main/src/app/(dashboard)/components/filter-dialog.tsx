import { SearchParams } from "@/api/events";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group";
import moment from "moment";
import { useEffect, useState } from "react";

export interface FilterDialogProps {
  open: boolean;
  onOpenChange: (value: boolean) => void;

  filters: SearchParams;
  onFiltersChange: (value: SearchParams) => void;
}

export default function FilterDialog(props: FilterDialogProps) {
  const [filtersCopy, setFiltersCopy] = useState<SearchParams>({});

  function clear() {
    setFiltersCopy({
      end_date: "",
      start_date: "",
    });
  }

  function save() {
    const odata = toOuter(filtersCopy);
    props.onFiltersChange(odata);
    props.onOpenChange(false);
  }

  useEffect(() => {
    if (props.open) {
      const data = toInternal(props.filters);
      const copied = JSON.parse(JSON.stringify(data));
      setFiltersCopy(copied);
    }
  }, [props.open, props.filters]);

  return (
    <Dialog open={props.open} onOpenChange={props.onOpenChange}>
      <DialogContent className="w-full md:min-w-[700px]">
        <DialogHeader>
          <DialogTitle>Фильтрация</DialogTitle>
          <DialogDescription>Настройки отображения</DialogDescription>
        </DialogHeader>
        <div className="grid grid-cols-1 gap-6 md:grid-cols-2">
          <div className="flex flex-col gap-2">
            <Label className="" htmlFor="filter-date-start">
              Отображать мероприятия с
            </Label>
            <div className="flex gap-2">
              <Input
                id="filter-date-start flex-grow"
                type="date"
                value={filtersCopy.start_date}
                onChange={(e) => {
                  filtersCopy.start_date = e.target.value;
                  setFiltersCopy({ ...filtersCopy });
                }}
              />
              <Button
                variant="secondary"
                onClick={() => {
                  filtersCopy.start_date = moment().format("YYYY-MM-DD");
                  setFiltersCopy({ ...filtersCopy });
                }}
              >
                Сегодня
              </Button>
            </div>
          </div>
          <div className="flex flex-col gap-2">
            <Label className="" htmlFor="filter-date-end">
              Отображать мероприятия по
            </Label>
            <div className="flex gap-2">
              <Input
                id="filter-date-end flex-grow"
                type="date"
                value={filtersCopy.end_date}
                onChange={(e) => {
                  filtersCopy.end_date = e.target.value;
                  setFiltersCopy({ ...filtersCopy });
                }}
              />
              <Button
                variant="secondary"
                onClick={() => {
                  filtersCopy.end_date = moment().format("YYYY-MM-DD");
                  setFiltersCopy({ ...filtersCopy });
                }}
              >
                Сегодня
              </Button>
            </div>
          </div>
          <div className="flex flex-col gap-2">
            <Label className="" htmlFor="filter-age-min">
              Минимальный возраст
            </Label>
            <div className="flex gap-2">
              <Input
                id="filter-age-min flex-grow"
                type="number"
                min="5"
                max="100"
                value={filtersCopy.min_age ?? ""}
                onChange={(e) => {
                  filtersCopy.min_age = e.target.value
                    ? parseInt(e.target.value)
                    : undefined;
                  setFiltersCopy({ ...filtersCopy });
                }}
              />
            </div>
          </div>
          <div className="flex flex-col gap-2">
            <Label className="" htmlFor="filter-age-max">
              Максимальный возраст
            </Label>
            <div className="flex gap-2">
              <Input
                id="filter-age-max flex-grow"
                type="number"
                min="5"
                max="100"
                value={filtersCopy.max_age ?? ""}
                onChange={(e) => {
                  filtersCopy.max_age = e.target.value
                    ? parseInt(e.target.value)
                    : undefined;
                  setFiltersCopy({ ...filtersCopy });
                }}
              />
            </div>
          </div>

          <div className="flex flex-col gap-2">
            <Label className="" htmlFor="filter-part-min">
              Минимально участников
            </Label>
            <div className="flex gap-2">
              <Input
                id="filter-part-min flex-grow"
                type="number"
                min="1"
                value={filtersCopy.min_participants ?? ""}
                onChange={(e) => {
                  filtersCopy.min_participants = e.target.value
                    ? parseInt(e.target.value)
                    : undefined;
                  setFiltersCopy({ ...filtersCopy });
                }}
              />
            </div>
          </div>
          <div className="flex flex-col gap-2">
            <Label className="" htmlFor="filter-part-max">
              Максимально участников
            </Label>
            <div className="flex gap-2">
              <Input
                id="filter-part-max flex-grow"
                type="number"
                min="1"
                value={filtersCopy.max_participants ?? ""}
                onChange={(e) => {
                  filtersCopy.max_participants = e.target.value
                    ? parseInt(e.target.value)
                    : undefined;
                  setFiltersCopy({ ...filtersCopy });
                }}
              />
            </div>
          </div>
          <div className="flex items-center gap-2 md:col-span-2">
            <Label className="" htmlFor="filter-sex">
              Пол:
            </Label>
            <div className="flex gap-2">
              <ToggleGroup
                type="single"
                id="filter-sex"
                value={
                  filtersCopy.sex != undefined
                    ? filtersCopy.sex
                      ? "man"
                      : "woman"
                    : "no"
                }
                onValueChange={(value) => {
                  if (value == "") return;
                  if (value == "no") filtersCopy.sex = undefined;
                  else filtersCopy.sex = value == "man";
                  setFiltersCopy({ ...filtersCopy });
                }}
              >
                <ToggleGroupItem value="no">Неважен</ToggleGroupItem>
                <ToggleGroupItem value="man">М</ToggleGroupItem>
                <ToggleGroupItem value="woman">Ж</ToggleGroupItem>
              </ToggleGroup>
            </div>
          </div>
          <div className="flex items-center gap-2 md:col-span-2">
            <Label className="" htmlFor="filter-loc">
              Местоположение:
            </Label>
            <div className="flex flex-grow gap-2">
              <Input
                id="filter-loc"
                value={filtersCopy.location ?? ""}
                onChange={(e) => {
                  filtersCopy.location =
                    e.target.value != "" ? e.target.value : undefined;
                  setFiltersCopy({ ...filtersCopy });
                }}
              />
            </div>
          </div>
        </div>
        <DialogFooter className="flex gap-2 pt-4 sm:justify-between">
          <div>
            <Button variant="secondary" onClick={() => clear()}>
              Сбросить
            </Button>
          </div>
          <div className="space-x-2">
            <Button
              variant="secondary"
              onClick={() => props.onOpenChange(false)}
            >
              Отмена
            </Button>
            <Button onClick={() => save()}>Сохранить</Button>
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

function toInternal(params: SearchParams): SearchParams {
  return {
    ...params,
    start_date: params.start_date ? dateFrom(params.start_date) : "",
    end_date: params.end_date ? dateFrom(params.end_date) : "",
  };
}

function toOuter(params: SearchParams): SearchParams {
  const data = {
    ...params,
    start_date: params.start_date ? dateTo(params.start_date) : "",
    end_date: params.end_date ? dateTo(params.end_date) : "",
  };
  return replaceEmptyStrings(data);
}

function dateFrom(dateS: string) {
  return moment(dateS, "DD.MM.YYYY").format("YYYY-MM-DD");
}

function dateTo(dateS: string) {
  return moment(dateS, "YYYY-MM-DD").format("DD.MM.YYYY");
}

function replaceEmptyStrings<
  T extends {
    [k: string]: any;
  },
>(
  obj: T
): {
  [k: string]: any;
} {
  return Object.fromEntries(
    Object.entries(obj).filter(([key, value]) => value !== "")
  );
}
