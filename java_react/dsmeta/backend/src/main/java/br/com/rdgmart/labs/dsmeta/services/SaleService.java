package br.com.rdgmart.labs.dsmeta.services;

import java.time.LocalDate;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import br.com.rdgmart.labs.dsmeta.entities.Sale;
import br.com.rdgmart.labs.dsmeta.repositories.SaleRepository;

@Service
public class SaleService {
	
	@Autowired
	private SaleRepository repository;
	
	
	public Page<Sale> findSales(String minDate, String maxDate, Pageable pageable) {
		LocalDate min = LocalDate.parse(minDate);
		LocalDate max = LocalDate.parse(maxDate);
		
		return repository.findSales(min, max, pageable);
	}

}
