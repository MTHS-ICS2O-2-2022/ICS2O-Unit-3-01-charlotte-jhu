// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: March 2023
// This file contains the JS functions for index.html

'use strict'

/**
 * This function calculates the are of a trapezoid
 */

function myButtonClicked() {
  // input
  const aBase = parseInt(document.getElementById("aBase").value)
  const bBase = parseInt(document.getElementById("bBase").value)
  const height = parseInt(document.getElementById("height").value)

  // process
  const area = [(aBase + bBase) / 2] * height

  // output
  document.getElementById('area').innerHTML = 'The area is ' + area + 'mm²'
}
