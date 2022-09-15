package br.com.rdgmart.labs.dsmeta.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import com.twilio.Twilio;
import com.twilio.rest.api.v2010.account.Message;
import com.twilio.type.PhoneNumber;

import br.com.rdgmart.labs.dsmeta.entities.Sale;
import br.com.rdgmart.labs.dsmeta.repositories.SaleRepository;

@Service
public class SmsService {

	@Value("${twilio.sid}")
	private String twilioSid;

	@Value("${twilio.key}")
	private String twilioKey;

	@Value("${twilio.phone.from}")
	private String twilioPhoneFrom;

	@Value("${twilio.phone.to}")
	private String twilioPhoneTo;

	@Autowired
	private SaleRepository saleRepository;
	
	public void sendSms(Long saleId) {
		
		if(saleId != null) {
			Sale sale = saleRepository.findById(saleId).get();
			
			if(sale != null) {
				
				String msg = "The Saller " + sale.getSellerName() + " sold " + String.format("%.2f", sale.getAmount());
				
				Twilio.init(twilioSid, twilioKey);
				
				PhoneNumber to = new PhoneNumber(twilioPhoneTo);
				PhoneNumber from = new PhoneNumber(twilioPhoneFrom);
				
				Message message = Message.creator(to, from, msg).create();
				
				System.out.println(message.getSid());
				return;
			}
			
		}
		
		System.out.println("Not send sms message ;( ");

	}
}
