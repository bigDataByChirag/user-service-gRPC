package logger

import (
	"log/slog"
	"os"
)

func SetupSlog() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		//AddSource: true: This will cause the source dir and line number of the logger message to be included in the output
		AddSource: true,
	})

	logger := slog.New(logHandler)
	//SetDefault makes l the default Logger. in our case we would be doing structured logging
	slog.SetDefault(logger)
}
