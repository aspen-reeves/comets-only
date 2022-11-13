import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
})
export class SignupComponent implements OnInit {
  signup!: FormGroup;

  constructor() {}

  ngOnInit(): void {
    this.signup = new FormGroup({
      name: new FormControl(null, [Validators.required]),
      age: new FormControl(null, [Validators.required]),
      lang: new FormControl(null, [Validators.required]),
      os: new FormControl(null, [Validators.required]),
      editor: new FormControl(null, [Validators.required]),
      lastShower: new FormControl(null, [Validators.required]),
      code: new FormControl(null, [Validators.required]),
    });
    console.log(this.signup);
  }

  onSubmit() {
    console.log(this.signup.value);
  }
}
