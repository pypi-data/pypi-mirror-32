//
//  ViperController.swift
//
//  Created by Bohdan Ivanov on 19.04.18.
//  Copyright © 2018 bivanov. All rights reserved.
//

import UIKit

class ViperController<T>: {{ base_viewcontroller }} where T: AbstractPresenter {
    var presenter: T!
}
