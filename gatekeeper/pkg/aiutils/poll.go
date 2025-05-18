package aiutils

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"time"
)

func PollRunStatus(ctx context.Context, client *openai.Client, threadID, runID string) (*openai.Run, error) {
	interval := 2

	for {
		run, err := client.RetrieveRun(ctx, threadID, runID)
		if err != nil {
			return nil, fmt.Errorf("failed to get run status: %v", err)
		}

		switch run.Status {
		case openai.RunStatusCompleted:
			return &run, nil
		case openai.RunStatusFailed, openai.RunStatusCancelled:
			return nil, fmt.Errorf("run failed or was cancelled: %s", run.Status)
		}

		time.Sleep(time.Second * time.Duration(interval))

		if interval < 5 {
			interval += 1
		}
	}
}
