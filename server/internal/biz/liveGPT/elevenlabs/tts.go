package elevenlabs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"faceto-ai/internal/biz/liveGPT/elevenlabs/types"
)

func (c Client) TTSWriter(ctx context.Context, w io.Writer, text, modelID, voiceID string, options types.SynthesisOptions) error {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s", voiceID)
	opts := types.TTS{
		Text:          text,
		ModelID:       modelID,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/taigrr/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	switch res.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200:
		defer res.Body.Close()
		io.Copy(w, res.Body)
		return nil
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = errors.Join(err, jerr)
		} else {
			err = errors.Join(err, ve)
		}
		return err
	}
}

func (c Client) TTS(ctx context.Context, text, voiceID, modelID string, options types.SynthesisOptions) ([]byte, error) {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s", voiceID)
	client := &http.Client{}
	opts := types.TTS{
		Text: text,
		//ModelID:       modelID,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	c.log.WithContext(ctx).Debugw("TTS.Client.Method", "TTS", "url", url, "options", opts)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		c.log.WithContext(ctx).Errorf("TTS.Client.Method.TTS http.NewRequestWithContext, err:%v", err)
		return []byte{}, err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/taigrr/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		c.log.WithContext(ctx).Errorf("TTS.Client.Method.TTS client.Do, err:%v", err)
		return []byte{}, err
	}
	c.log.WithContext(ctx).Debugf("TTS.Client.Method.TTS StatusCode:%d", res.StatusCode)
	switch res.StatusCode {
	case 401:
		return []byte{}, ErrUnauthorized
	case 200:
		b := bytes.Buffer{}

		defer res.Body.Close()
		io.Copy(&b, res.Body)
		return b.Bytes(), nil
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = errors.Join(err, jerr)
		} else {
			err = errors.Join(err, ve)
		}
		return []byte{}, err
	}
}

func (c Client) TTSStream(ctx context.Context, w io.Writer, text, voiceID string, options types.SynthesisOptions) error {
	options.Clamp()
	url := fmt.Sprintf(c.endpoint+"/v1/text-to-speech/%s/stream", voiceID)
	opts := types.TTS{
		Text:          text,
		VoiceSettings: options,
	}
	b, _ := json.Marshal(opts)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("xi-api-key", c.apiKey)
	req.Header.Set("User-Agent", "github.com/taigrr/elevenlabs")
	req.Header.Set("accept", "audio/mpeg")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200:
		defer res.Body.Close()
		io.Copy(w, res.Body)
		return nil
	case 422:
		fallthrough
	default:
		ve := types.ValidationError{}
		defer res.Body.Close()
		jerr := json.NewDecoder(res.Body).Decode(&ve)
		if jerr != nil {
			err = errors.Join(err, jerr)
		} else {
			err = errors.Join(err, ve)
		}
		return err
	}
}
