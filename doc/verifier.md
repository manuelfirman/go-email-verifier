# Verifier
Verificador de registros DNS relacionados con el correo electrónico (MX, SPF y DMARC) para un dominio dado.


### Variables

##### :ballot_box_with_check: hasMX (bool)
Indica si el dominio tiene **registros MX (Mail Exchange)**. Los registros MX especifican los servidores de correo electrónico responsables de recibir los correos electrónicos destinados a ese dominio. 
*Si hasMX es `true`, significa que el dominio tiene servidores de correo configurados.*


##### :ballot_box_with_check: hasSPF (bool)
Indica si el dominio tiene un **registro SPF (Sender Policy Framework)**. SPF es un mecanismo de autenticación de correo electrónico que ayuda a prevenir el correo electrónico no deseado falsificado. 
*Si hasSPF es `true`, significa que el dominio tiene una política SPF configurada.*
##### :ballot_box_with_check: hasDMARC (bool)
Indica si el dominio tiene un **registro DMARC (Domain-based Message Authentication, Reporting, and Conformance)**. DMARC es una política de autenticación de correo electrónico que permite a los propietarios de dominios proteger su dominio contra el uso indebido en los correos electrónicos falsificados. 
*Si hasDMARC es `true`, significa que el dominio tiene una política DMARC configurada.*
##### :ballot_box_with_check: spfRecord (string)
Almacena el **contenido del registro SPF del dominio**. El registro SPF especifica las reglas que indican qué servidores están autorizados para enviar correos electrónicos en nombre del dominio. 
*Si hasSPF es `true`, este campo contendrá el registro SPF del dominio.*
##### :ballot_box_with_check: dmarcRecord (string)
Almacena el **contenido del registro DMARC del dominio**. El registro DMARC especifica cómo deben procesar los receptores los correos electrónicos que supuestamente provienen del dominio. 
*Si hasDMARC es `true`, este campo contendrá el registro DMARC del dominio.*

### Uso
```bash
go run .\cmd\main.go
Enter the domain names to check, one per line. Press Ctrl+C to finish.
> invalid # ingreso de un dominio invalido
2024/03/22 15:39:31 Error looking up MX records for 'invalid': lookup invalid: dnsquery: DNS name does not exist.
--------------------------------------------------
> gmail.com # ingreso de un dominio valido
Domain: gmail.com
hasMX: true
hasSPF: true
hasDMARC: true
spfRecord: v=spf1 redirect=_spf.google.com
dmarcRecord: v=DMARC1; p=none; sp=quarantine; rua=mailto:mailauth-reports@google.com
--------------------------------------------------
```