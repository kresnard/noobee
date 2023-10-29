package main

import (
	"async_smtp/mail"
	"log"
)

func main() {
	// setup email tujuan
	to := []string{"kranggad@gmail.com"}

	// setup cc
	cc := []string{"sosmedkresna@gmail.com"}

	subject := "Tes Email!"
	// message := "Hello from first email"
	message := `
	<html>
		<body>
			<h1> Hello From Kresna</h1>
			<button class="btn btn-primary ">Click Me</button>
		</body>
	</html>
	`

	// panggil sungsi send email
	// err := mail.SendEmail(to, cc, subject, message)
	err := mail.SendMailGomail(to, cc, subject, message)
	if err != nil {
		panic(err)
	}

	log.Println("success send mail to", append(to, cc...))
}
