import re


def toggle_torch_source(file_path):
    try:
        with open(file_path, "r") as file:
            content = file.read()

        torch_pattern = (
            r'torch = \[\s*{version\s*=\s*"[^"]+",\s*source\s*=\s*"([^"]+)"}\s*\]'
        )
        match = re.search(torch_pattern, content)

        if not match:
            print("Не удалось найти запись для torch в pyproject.toml.")
            return

        current_source = match.group(1)

        sources = ["torchcpu", "torch_rocm", "torch_cuda"]
        if current_source not in sources:
            print(
                f"Неизвестный источник: {current_source}. Убедитесь, что значение корректно."
            )
            return

        new_source = sources[(sources.index(current_source) + 1) % len(sources)]

        updated_content = re.sub(
            torch_pattern,
            f'torch = [{{version = "^2.5.1", source="{new_source}"}}]',
            content,
        )

        with open(file_path, "w") as file:
            file.write(updated_content)

        print(f"Источник torch успешно изменен на {new_source}.")
    except Exception as e:
        print(f"Произошла ошибка: {e}")


# Использование
file_path = "pyproject.toml"
toggle_torch_source(file_path)
