package constants

const (
	MailTemplate = `
	<!DOCTYPE html>
<head>
    <meta content="IE=edge,chrome=1" http-equiv="X-UA-Compatible">
    <meta charset="utf-8">
    <meta content="text/html" http-equiv="Content-Type">
    <meta name="viewport" content="width=device-width, initial-scale=2">
</head>
<html>
    <body style="background-color:#fff; padding:0 0 15px; margin:15px auto; color:#999999; font-size:14px;width:auto;overflow-x: hidden !important;">
        <link href='https://fonts.googleapis.com/css?family=Montserrat:600' rel='stylesheet' type='text/css'>
        <div style="font-family: 'Montserrat', sans-serif !important;padding:20px 5%; border-bottom:solid 1px #C1CAE3;">
            <a href="#" style="float:left">
                <link style="width:130px;" src="/images/favicon.ico" alt="">
            </a>
            <div style="clear:both;"></div>
        </div>
        <div style="clear:both;"></div>
        <div style="font-family: 'Montserrat', sans-serif !important; margin:0 5%; line-height:150%;">
            <div style="font-size:22px; margin:25px auto;max-width: 700px; line-height:160%; color: #4986f2;text-align: center;">
                
                Hai,<br />
                Anda mendapatkan notifikasi sebagai berikut :<br />
                
            </div>
            <div>
                <p style="color:#333; margin: 40px 0 0;font-size: 16px">{{.Message}}</p>
            </div>
            <div style="margin:20px 0;">
            </div>
            <img src='https://www.google-analytics.com/collect?v=1&tid=UA-891770-59&cid="/spD3tZiYTQwYBBeoRp6zqVHQuvGNZ0Ya5OfctjnIkw="dibalas&t=event&ec=email_notif_komentar&ea=email_dilihat_dibalas&el="/spD3tZiYTQwYBBeoRp6zqVHQuvGNZ0Ya5OfctjnIkw="&cs=email_notif_komentar&cm=email&cn=Email%20Notif%20Komentar&dp=email_dilihat/dibalas&dt=dibalas'/>
        </div>
        <div style="clear:both;"></div>
        <div style="font-family: 'Montserrat', sans-serif !important;text-align:center; margin:10% 0;"><br> <br> Anda menerima email ini karena Anda mengaktifkan notifikasi komentar.<br> Anda dapat mengatur dan <a href="#">UNSUBSCRIBE</a> notifikasi di halaman pengaturan profil anda </div>
    </body>
</html>
	`
)
