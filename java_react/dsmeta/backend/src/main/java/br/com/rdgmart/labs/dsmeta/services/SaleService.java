package br.com.rdgmart.labs.dsmeta.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import br.com.rdgmart.labs.dsmeta.entities.Sale;
import br.com.rdgmart.labs.dsmeta.repositories.SaleRepository;

@Service
public class SaleService {
	
	@Autowired
	private SaleRepository repository;
	
	
	public List<Sale> findSales() {
		return repository.findAll();
	}

}
