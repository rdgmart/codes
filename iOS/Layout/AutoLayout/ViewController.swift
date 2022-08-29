//
//  ViewController.swift
//  AutoLayout
//
//  Created by Rodrigo Martins on 22/08/22.
//

import UIKit

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
    }

    @IBAction func logout(_ sender: UIButton) {
        
        if let nav = navigationController{
            nav.popViewController(animated: true);
        }
        
        self.view.window?.rootViewController?.dismiss(animated: true, completion: nil)

        
    }
    
}

