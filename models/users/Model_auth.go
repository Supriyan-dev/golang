package users

import (
	"../../models"
	utils_enter_the_information "../../utils/enter_the_information"
	"crypto/md5"
	"encoding/hex"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Models_init_Users models.DB_init

func (model Models_init_Users) ForgotPasswordWithEmail(email string) (responseEmail int, err error) {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	yearTokyo := time.Now().In(loc).Year()

	var checkEmailInt int
	checkEmail := model.DB.QueryRow(`select count(*) from user where email = ?`, email).Scan(&checkEmailInt)

	var DataEmployee string
	GetDataEmploye := model.DB.QueryRow(`select employee_number from user where email =?`, email).Scan(&DataEmployee)

	if GetDataEmploye != nil {
		log.Println(GetDataEmploye)
	}

	if checkEmail != nil {
		log.Println(checkEmail.Error())
	}

	if checkEmailInt > 0 {
		responseEmail = 1
	} else {
		responseEmail = 0
	}

	//send email
	if checkEmailInt > 0 {

		RandomInte := rand.Intn(999999)
		var RandomInt int

		checkIntRandom := utils_enter_the_information.CheckDataByIdInt(`select COUNT(*) from user where recovery_pin = ?`, RandomInte)

		if checkIntRandom == 0 {
			RandomInt = RandomInte
		} else {
			for {
				RandomInteg := rand.Intn(999999)
				checkIntRandom := utils_enter_the_information.CheckDataByIdInt(`select COUNT(*) from user where recovery_pin = ?`, RandomInteg)
				if checkIntRandom == 0 {
					RandomInt = RandomInteg
					break
				}
			}
		}

		ForgotAction, errForgotAction := model.DB.Exec(`update user set recovery_pin = ? where email =?`, RandomInt, email)

		if errForgotAction != nil {
			log.Println(errForgotAction)
		}

		rowsAffected, _ := ForgotAction.RowsAffected()

		if rowsAffected == 1 {

			htmlContext := `<!doctype html>
<html>

<head>
	<meta name="viewport" content="width=device-width" />
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<title>Forgot Password</title>
	<link href="https://fonts.googleapis.com/css?family=Muli&display=swap" rel="stylesheet">
	<link href="https://fonts.googleapis.com/css?family=Montserrat&display=swap" rel="stylesheet">
	<style>
		img {
			border: none;
			-ms-interpolation-mode: bicubic;
			max-width: 100%;
		}

		body {
			background-color: #f6f6f6;
			font-family: 'Montserrat', sans-serif !important;
			-webkit-font-smoothing: antialiased;
			font-size: 14px;
			line-height: 1.4;
			margin: 0;
			padding: 0;
			-ms-text-size-adjust: 100%;
			-webkit-text-size-adjust: 100%;
		}

		.font-sans {
			font-family: sans-serif !important;
		}

		table {
			border-collapse: separate;
			mso-table-lspace: 0pt;
			mso-table-rspace: 0pt;
			width: 100%;
		}

		table td {
			font-family: 'Muli', sans-serif !important;
			font-size: 14px;
			vertical-align: top;
			padding: 3px 0px;
		}

		.body {
			background-color: #f6f6f6;
			width: 100%;
		}

		.container {
			display: block;
			margin: 0 auto !important;
			max-width: 580px;
			padding: 10px;
			width: 580px;
		}

		.content {
			box-sizing: border-box;
			display: block;
			margin: 0 auto;
			max-width: 580px;
			padding: 10px;
		}

		.main {
			background: #ffffff;
			border-radius: 3px;
			width: 100%;
		}

		.wrapper {
			box-sizing: border-box;
			padding: 20px;
			border: 1px solid #ddd;
			/* border: 1px solid #52c234; */
		}

		.content-block {
			padding-bottom: 10px;
			padding-top: 10px;
		}

		.footer {
			clear: both;
			margin-top: 10px;
			text-align: center;
			width: 100%;
		}

		.footer td,
		.footer p,
		.footer span,
		.footer a {
			color: #999999;
			font-size: 12px;
			text-align: center;
		}

		h1,
		h2,
		h3,
		h4 {
			color: #000000;
			font-family: 'Muli', sans-serif;
			font-weight: 400;
			line-height: 1.4;
			margin: 0;
			margin-bottom: 30px;
		}

		h1 {
			font-size: 35px;
			font-weight: 300;
			text-align: center;
			text-transform: capitalize;
		}

		p,
		ul,
		ol {
			font-family: 'Muli', sans-serif;
			font-size: 14px;
			font-weight: normal;
			margin: 0;
			margin-bottom: 15px;
		}

		p li,
		ul li,
		ol li {
			list-style-position: inside;
			margin-left: 5px;
		}

		a {
			color: #52c234;
			text-decoration: underline;
		}

		.btn {
			box-sizing: border-box;
			width: 100%;
		}

		.btn>tbody>tr>td {
			padding-bottom: 15px;
		}

		.btn table {
			width: auto;
		}

		.btn table td {
			background-color: #ffffff;
			border-radius: 5px;
			text-align: center;
		}

		.btn a {
			background-color: #52c234;
			border: solid 1px #52c234;
			border-radius: 5px;
			box-sizing: border-box;
			color: #52c234;
			cursor: pointer;
			display: inline-block;
			font-size: 14px;
			font-weight: bold;
			margin: 0;
			padding: 12px 25px;
			text-decoration: none;
			text-transform: capitalize;
		}

		.btn-primary table td {
			background-color: #52c234;
		}

		.btn-primary a {
			background-color: #52c234;
			border-color: #52c234;
			color: #ffffff;
		}

		.last {
			margin-bottom: 0;
		}

		.first {
			margin-top: 0;
		}

		.align-center {
			text-align: center;
		}

		.align-right {
			text-align: right;
		}

		.align-left {
			text-align: left;
		}

		.clear {
			clear: both;
		}

		.mt0 {
			margin-top: 0;
		}

		.mb0 {
			margin-bottom: 0;
		}

		.preheader {
			color: transparent;
			display: none;
			height: 0;
			max-height: 0;
			max-width: 0;
			opacity: 0;
			overflow: hidden;
			mso-hide: all;
			visibility: hidden;
			width: 0;
		}

		.powered-by a {
			text-decoration: none;
		}

		hr {
			border: 0;
			/* border-bottom: 1px solid #52c234; */
			border-bottom: 1px solid #eee;
			margin: 20px 0;
		}

		@media only screen and (max-width: 620px) {
			table[class=body] h1 {
				font-size: 28px !important;
				margin-bottom: 10px !important;
			}

			table[class=body] p,
			table[class=body] ul,
			table[class=body] ol,
			table[class=body] td,
			table[class=body] span,
			table[class=body] a {
				font-size: 16px !important;
			}

			table[class=body] .wrapper,
			table[class=body] .article {
				padding: 10px !important;
			}

			table[class=body] .content {
				padding: 0 !important;
			}

			table[class=body] .container {
				padding: 0 !important;
				width: 100% !important;
			}

			table[class=body] .main {
				border-left-width: 0 !important;
				border-radius: 0 !important;
				border-right-width: 0 !important;
			}

			table[class=body] .btn table {
				width: 100% !important;
			}

			table[class=body] .btn a {
				width: 100% !important;
			}

			table[class=body] .img-responsive {
				height: auto !important;
				max-width: 100% !important;
				width: auto !important;
			}
		}

		@media all {
			.ExternalClass {
				width: 100%;
			}

			.ExternalClass,
			.ExternalClass p,
			.ExternalClass span,
			.ExternalClass font,
			.ExternalClass td,
			.ExternalClass div {
				line-height: 100%;
			}

			.apple-link a {
				color: inherit !important;
				font-family: inherit !important;
				font-size: inherit !important;
				font-weight: inherit !important;
				line-height: inherit !important;
				text-decoration: none !important;
			}

			#MessageViewBody a {
				color: inherit;
				text-decoration: none;
				font-size: inherit;
				font-family: inherit;
				font-weight: inherit;
				line-height: inherit;
			}

			.btn-primary table td:hover {
				background-color: #34495e !important;
			}

			.btn-primary a:hover {
				background-color: #34495e !important;
				border-color: #34495e !important;
			}
		}

		.FontHarga {
			color: green;
			font-size: 22pt;
		}

		.FontInvoice {
			color: green;
			font-size: 22pt;
		}

		.FontVirtual {
			/* color:orange; */
			color: green;
			font-size: 18pt;
		}

		.FontBatas {
			color: #C90000;
			font-size: 12pt;
		}

		.quo {
			color: #FF0000 !important;
		}

		.muted {
			color: #444;
		}
	</style>
</head>

<body class="">
	<table role="presentation" border="0" cellpadding="0" cellspacing="0" class="body">
		<tr>
			<td>&nbsp;</td>
			<td class="container">
				<div class="content">
					<table role="presentation" class="main">
						<tr>
							<td class="wrapper">
								<table role="presentation" border="0" cellpadding="0" cellspacing="0">
									<tr>
										<td>
											<center>
												<img src="https://ci6.googleusercontent.com/proxy/uMyhwlZ9jCzz-aP1cTrDvcz59rw-eJ211d3uLAiumei7i7Ozc0Wzn_BJiuL97f_ldlJ_2vmqgcby9VFtJWh3OeA43A=s0-d-e1-ft#http://104.41.162.39/images/Icon/Kasumi_Logo.gif" alt="" width="180">
											</center>
											<hr>

											<h3 style="text-align:center;font-family: 'Montserrat', sans-serif;color:#777;"> <span class="quo" style="color:#FF0000!important;font-weight:bold;">[ `+DataEmployee+` ]</span>
												<br>
												パスワードの変更が要求されました。下記の６桁のパスコードを入力してください。
												<br>
												<center style="margin-top:2%!important;">
													<h3 style="font-size:14pt;color:#FF0000!important;font-weight:bold;">
														`+strconv.Itoa(RandomInt)+`
													</h3>
												</center>
											</h3>
											<table>
												<tbody class="muted">
													<tr>
														<td colspan="2" align="center" class="quo">
															<br>
															<small style="font-size:7pt;text-align:center;color:#FF0000!important;">
																kasumi.co.jp
															</small>
															<small style="font-size:7pt;text-align:center;color:#888;padding-top:10px;border-bottom:#eee solid 1px;">
																<br>
																All Rights Reserved. &copy;  . `+strconv.Itoa(yearTokyo)+`
															</small>
														</td>
													</tr>

												</tbody>
											</table>

											<br>
										</td>
									</tr>
								</table>
							</td>
						</tr>

					</table>
					<div class="footer">
						<table role="presentation" border="0" cellpadding="0" cellspacing="0">
							<tr>
								<td class="content-block">
									<span class="apple-link">

									</span>.
								</td>
							</tr>
							<tr>
								<td class="content-block powered-by">
								</td>
							</tr>
						</table>
					</div>

				</div>
			</td>
			<td>&nbsp;</td>
		</tr>
	</table>
</body>

</html>`

			log.Println(rowsAffected)
			//hasher := md5.New()
			//hasher.Write([]byte(email))
			//tokenString := hex.EncodeToString(hasher.Sum(nil))

			//linkForgotPassword := `link/` + email + `/` + tokenString
			from := mail.NewEmail("Workflow Kasumi", "siapasayachannel@gmail.com")
			subject := "パスワードをお忘れですか - Workflow Kasumi"
			to := mail.NewEmail("", email)
			plainTextContent := `hello` + email
			htmlContent := htmlContext
			message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
			client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
			_, err := client.Send(message)
			if err != nil {
				log.Println(err)
			} else {
				//fmt.Println(response.StatusCode)
				//fmt.Println(response.Body)
				//fmt.Println(response.Headers)
				return responseEmail, nil
			}
		}

	}
	return responseEmail, nil
}

func (model Models_init_Users) ForgotPasswordWithPin(pin string) (responseForgot string, err error) {
	var checkPinInt int
	checkPin := model.DB.QueryRow(`select count(*) from user where recovery_pin = ?`, pin).Scan(&checkPinInt)

	if checkPin != nil {
		log.Println(checkPin)
	}

	if checkPinInt > 0 {
		return "Success Response", nil
	}

	return "Bad Request", nil
}

func (model Models_init_Users) ForgotPasswordAction(email string, pin string, NewPassword string) (responseForgot string, err error) {
	var checkPinInt int
	checkPin := model.DB.QueryRow(`select count(*) from user where recovery_pin = ?`, pin).Scan(&checkPinInt)

	if checkPin != nil {
		log.Println(checkPin)
	}
	log.Println(checkPinInt)
	if checkPinInt > 0 {

		hasher := md5.New()
		hasher.Write([]byte(NewPassword))
		NewPasswordString := hex.EncodeToString(hasher.Sum(nil))

		ForgotAction, errForgotAction := model.DB.Exec(`update user set password = ? where recovery_pin =? and email =?`, NewPasswordString, pin, email)

		if errForgotAction != nil {
			log.Println(errForgotAction)
		}

		rowsAffected, _ := ForgotAction.RowsAffected()

		if rowsAffected == 1 {
			return "Success Response", nil
		}
	}

	return "Bad Request", nil
}
