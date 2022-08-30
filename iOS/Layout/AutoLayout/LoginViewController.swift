//
//  LoginViewController.swift
//  AutoLayout
//
//  Created by Rodrigo Martins on 23/08/22.
//

import Foundation
import UIKit

class LoginViewController: UIViewController {
    
    
    @IBOutlet weak var textLogin: UITextField!
    
    @IBOutlet weak var textPwd: UITextField!
    
    override func viewDidLoad() {
        super.viewDidLoad()
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        if segue.identifier == "afterLogin"{
            
            if let vcDestination = segue.destination as? ViewController{

                if let text = textLogin.text{
                    vcDestination.loginUser = text
                }else {
                    vcDestination.loginUser = ""
                }
                
            }
        }
    }
}


