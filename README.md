# Valida CPF Backend

It is a Web application in GoLang that receives a CPF (Brazilian Tax Identification Document) as an argument and returns whether it is a valid CPF, and also formats the numbers in the pattern with dots and dash.

## **Why is this important?**

The CPF is a Brazilian Tax Identification Document, so the number must be correct before we can send it for a credit review or government verification.
Either because a customer is checking credit or even just filling out a registration form.
This way we can ensure that there are no errors before we overload the system with the next step.

## **How it works?**

In the CPF, the last two numbers are verification digits. The calculation works starting with weights associated with the first 9 digits, and a division by the number 11. Then the module is subtracted from 11 to find the first digit. After this step, considering the tenth digit we found, we redo the calculation to find the second check digit. That is, with the first 9 numbers we can identify whether the CPF is correct or not.

## **How to use?**

It's very easy, with a `go run .` you will serve port `8080`, so in the local host type: http://localhost:8080/valida-cpf?numero=91468384066

These eleven numbers are a CPF, which our application will check if it is correct. You can replace this number with another one for testing.

If the number is a correct CPF you will receive the response `{"valido":true,"formatado":"914.683.840-66"}`, otherwise it is not a correct number you will receive a `"valido":"false"` and the number formatted.
