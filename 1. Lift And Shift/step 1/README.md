# Update Private IPs of the instances to Route53 private DNS ZONES

## To create Hosted zone

1. Go to Route53 and select hosted zone
![Capture9](https://user-images.githubusercontent.com/18073289/216613692-d6c63b8b-9706-4e3f-8b23-59d538185c6d.PNG)

2. Type in domain name(sam.org), Select the Region, Select default vpc and Create hosted zone
![Capture10](https://user-images.githubusercontent.com/18073289/216613728-c42aa0c1-6b14-44fb-a0c9-4cfb4ca76cae.PNG)

## To create records

## Create Simple Records
 1. Select simple routing
 <br />
 ![Capture11](https://user-images.githubusercontent.com/18073289/216613849-7e1e3b50-7cbf-4876-920e-1ac2759843ab.PNG)
 <br />
 <br />


 2. Type the sub domain (db01, mc01,bmq01) - Enter ip address of the corresponding ec2 instance
 <br />
 ![Capture10](https://user-images.githubusercontent.com/18073289/216614146-33d013b3-0397-47d0-97a6-3c7324a0d79d.PNG)
![Capture11](https://user-images.githubusercontent.com/18073289/216614208-2cc1a60c-579b-445a-aafa-f4fabac2b92f.PNG)
