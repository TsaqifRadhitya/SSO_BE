package Provider

import (
	"SSO_BE_API/Config"
	"context"
	"fmt"

	"github.com/resend/resend-go/v2"
)

type Client struct {
	*resend.Client
}

func InitClientMailer() *Client {
	return &Client{
		Client: resend.NewClient(Config.RESEND_API_KEY),
	}
}

func (c *Client) SendOtp(toEmail string, otp string) error {
	params := &resend.SendEmailRequest{
		From:    Config.RESEND_DOMAIN,
		To:      []string{toEmail},
		Subject: "Your OTP Code",
		Text:    fmt.Sprintf("Your OTP code is: %s", otp),
	}

	_, err := c.Emails.SendWithContext(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}
	return nil
}

func (c *Client) SendResetPasswordOTP(toEmail string, otp string) error {
	params := &resend.SendEmailRequest{
		From:    Config.RESEND_DOMAIN,
		To:      []string{toEmail},
		Subject: "Your Reset Password Code",
		Text:    fmt.Sprintf("Your OTP code is: %s", otp),
	}

	_, err := c.Emails.SendWithContext(context.Background(), params)

	return err
}
