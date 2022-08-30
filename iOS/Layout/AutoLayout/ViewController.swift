//
//  ViewController.swift
//  AutoLayout
//
//  Created by Rodrigo Martins on 22/08/22.
//

import UIKit

class ViewController: UIViewController {

    @IBOutlet weak var labelWelcome: UILabel!
    var loginUser:String = ""
        
    override func viewDidLoad() {
        super.viewDidLoad()
            
        if(loginUser != ""){
            labelWelcome.text = "Welcome " + loginUser
        }else{
            labelWelcome.text = "Hello unknown user"
        }
    }

    @IBAction func logout(_ sender: UIButton) {
        
        if let nav = navigationController{
            nav.popViewController(animated: true);
        }
        
        self.view.window?.rootViewController?.dismiss(animated: true, completion: nil)

        
    }
    
}

