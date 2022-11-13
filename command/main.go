package main

type Command interface {
	Execute() error
}

type Receiver interface {
	Action() error
}

type Invoker interface {
	Invoke() error
}

var _ Command = (*SendEmail)(nil)

type SendEmail struct {
	receiver Receiver
}

func (s *SendEmail) Execute() error {
	return s.receiver.Action()
}

var _ Receiver = (*SesClient)(nil)

type SesClient struct {
	from string
}

func (s *SesClient) Action() error {
	println("send email by ses")
	return nil
}

var _ Receiver = (*SendGrid)(nil)

type SendGrid struct {
	from string
}

func (s *SendGrid) Action() error {
	println("send email by sedngrid")
	return nil
}

var _ Receiver = (*Pinpoint)(nil)

type Pinpoint struct{}

func (p *Pinpoint) Action() error {
	println("send push notification by pinpoint")
	return nil
}

var _ Command = (*SendPushNotification)(nil)

type SendPushNotification struct {
	receiver Receiver
}

func (p *SendPushNotification) Execute() error {
	return p.receiver.Action()
}

var _ Invoker = (*Notify)(nil)

type Notify struct {
	commands []Command
}

func (n *Notify) Invoke() error {
	for _, c := range n.commands {
		if err := c.Execute(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	notify := &Notify{}
	notify.commands = append(
		notify.commands,
		&SendEmail{receiver: &SesClient{from: "example.com"}},
		&SendPushNotification{receiver: &Pinpoint{}},
	)
	notify.Invoke()
}
