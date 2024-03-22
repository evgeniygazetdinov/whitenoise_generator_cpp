package verifyemail

import "work_in_que/integrations/sendgrid"

func SendVerifyEmail(fullName string, email string, validationCode string) error {
	plainTextContent := "Hi " + fullName + ",\n\nWelcome to my app! We are very excited to have you use it. Confirming your email is the final step in your sign up process. Please use this link to confirm: https://yourwebsite.com/complete-validation?code=" + validationCode + "&email=" + email + "\n\n Link expire? Request a new one here: https://yourwebsite.com/request-validation?email=" + email
	htmlContent := "<html> <body> <div style='width: 98%; margin-left:auto; margin-right:auto; padding-top: 15px; padding-bottom: 20px; background-color: #f1f1f1;'> <div style='background-color: #fff; width: 90%; margin-left:auto; margin-right:auto;padding-top: 15px;'> <div style='width: 100%;padding: 10px 20px;'> <img src='https://upload.wikimedia.org/wikipedia/commons/thumb/0/08/Circle-icons-rocket.svg/1200px-Circle-icons-rocket.svg.png' style='height: 50px;' /> <h1 style='font-size: 32px;font-family: sans-serif;margin-top: 0px; margin-bottom: 0px;padding-top: 10px; padding-bottom: 10px;'>Confirm Your Email</h1> </div> <div style='padding: 10px 25px;'> <p style='margin-top: 0px; margin-bottom: 0px; font-size: 16px; font-family: sans-serif;'> Hi " + fullName + ",<br /> <br /> Welcome to my app! We are very excited to have you use it. Confirming your email is the final step in your sign up process. Please use the button below to confirm: </p> <div style='width: 100%; padding: 20px 5px; text-align: center;'> <a href=\"https://yourwebsite.com/complete-validation?code=" + validationCode + "&email=" + email + "\" style='padding: 10px 50px; background-color: #0489B1; color: #fff; border-width: 0px;font-size: 20px; text-decoration: none; font-family: sans-serif;'>Confirm Email</a> </div> <p style='font-size: 16px; font-family: sans-serif; padding-top: 15px;'> Link expire? <a href=\"https://yourwebsite.com/request-validation?email=" + email + "\">Request a new one</a>. </p> </div> <div style='padding-top: 15px; padding-bottom: 25px; text-align: center;'> <p style='font-size: 14px; font-family: sans-serif; margin-top: 0px; margin-bottom: 0px;'>Made by KeithWeaver</p> <p style='font-size: 12px; font-family: sans-serif; margin-top: 0px; margin-bottom: 0px; padding: 10px 0px;'> <a href='https://yourwebsite.com/blog' style='text-decoration: underline; color: #2e2e2e;'> Our Blog </a> <a href='https://yourwebsite.com/privacy' style='text-decoration: underline; color: #2e2e2e; padding: 0px 15px;'> Our Privacy Policy </a> <p> </div> </div> </div> </body> </html>"
	return sendgrid.SendEmail(fullName, email, "Confirm Email", plainTextContent, htmlContent)
}
