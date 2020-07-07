package coins

// DAI coinfactory information
var DAI = Coin{
	Info: CoinInfo{
		Icon:         "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAMAAACdt4HsAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAMAUExURUxpcb+1I/KaMfm7OP+/MfrHU///Sv/xFfuPFvirE/8AAPmsFPmuGvvDSvitF/rFT/m4MPm2K/qwH/nBRfmyJPzLVfvGUPmzHPm0J/q9PPmuGv+2SPmtGvzGSvuxG/u1H/m/QvvCRPmxHPW/TvapEfqvH/qtEvm/P/nBQ/vIU/3MVvqtEfm5M/vJV/zLVvq9PPzIT/mwG/qwG/yzF/zFTPqwGvvDSfrGTfSrFP3MUvrBRvm8QvqwFvzLU/zFRvzLUvnATvq0I/qxIvy2I/nBR/e/QvuuFPq0JPm/QfywFPbBV/m+QPq0JfuuEvm3MPzNVvqrEP+vE/mrEfm7NvmyJvutEvuuEvvAPvu0J//DLPq1Kvu1Kfm7O/m2LPm6OPm3Lvq+Pfm2Lv/NV/zKV//MWvzKV/////mvG/q/QfvFTvrCSP/+/vu+PfmtFvu0JvqxH/vDS/qwHvvDSfqwHfvETP/++/vETfq4MPqyIvq1KvmsFPq3LvmuGvrCR/vCRPqxIPq/QPmuGPq5MvrBRfqyIfqzJPq2K/zGUPq6NPq1Kfq9Ovq2LPu7NfrAQ/rARPqzJfu5Mvq+P/q8Pfq0KPu7NvvGUPq8Ofq4Mfq8Ov/9+fq5MPvHUvzYjPvLZfvEUvq2KvqyIPq7OPqxIfqwG/q4LfvJXvzNb////vvHU//+/Pu7O//9+/q1JP/8+P/UU/+9Hv3JVv/WWfqzIfzLZ/qrEfzLY/7JU/y3LPvHWP/FPfu9QfvET/vCTfvKYfutEvvIVPzUgvvIW//PSv/AJv+5FfqvGPvFVPvNavvBSfuuFfqxHf++LPq/Q/vRdvzakfzWhv3dmfu/RP/68P/56/7x1v/RUf+6Gf/NRfzPcv+0E/+zEv7pvf/AOf/AMP/RWfq9PP/89/3krv3hpvq7N/7uzf+2GP+6IvywF/7FQ/26Mv+yFf66KP29N/+9Mf+vEv/NVvzTfP/IQ/+yFP/NWPzakv725f7rxP3fof714f/PWP/SWv715PzRev3ksPvGV/7syVyN5ZEAAAB/dFJOUwAEBv0E/QIBAv0B+/v6/vz9/Pv9/O5lLP38lQdn+sn6+/GNGhwY7f0rlY+T3fralLC5sPawsJW5MfZwF3LG/XIx73D9y43v8nx8Hd182PB83EO+8N3SYkFB/tSSYmK/ldS/Q74t0v////////////////////////////////6H86rYAAAGgklEQVRYw5WXd1wTZxzGX5JADgKGjQiKGwfDva177723XilaQGwKmCJpQjAQIVJaIClLkdI4wnKADNlTBRyIEzTuvau1tX3vEnJ3IUD4/pHP5ZfneXL33nvvvT8AWoEw4AfDbuXgpYv7G588adx/8dLBK+3wIgI6hgmVtCWDV7BuXL/xsuAkpOAlPGStGLyEBiOYHfmhve+cBddL3xYYY+yF4AcFb0uvL5jTFxe0gxG0955Z+tSYpfISGMPS09KZvfviorYwAMiE/qVVLNZenbBYVaX9JyBQphsTBhjTo+ERq8v3bdKF9aihxxjAMNE5eghwHP6gPbsq4sFwR4DoGEsEmHRrKDP9uUNMyxq6mQCktZ82Ms20y0960MU0bSRNO4GJ0IammX6rJ6ZpQ2nUq2AymMP092MJw5iUOcUA3dJM93QC07Ru5CnFAI5plrs7hWWaI5FgBOycnHZ7d4rdTk52LXPShGHQ46FlEMFhLYJ0Yfmwh4F6QhmAqSWWwQT+Qclk9gT5S4T+wa2wLJmqmtQIs9+UaYf9CYKjjpIRRSUXVAXBDC0OT5vSj4ngI9C7xFpI4JVcKZVKTzSfwGluvvTlv7pT4irvKKEW1iW9sVFAQL9ZXtYSgqjkPLQVH+oOVUkkQgkZa69Z/aCdAeaXWHuRCK3OS01NlWrA/Kkoeqmw2j9UQhZ6WZfMh3YmbfmzrlEkYAAqRy/8Wn9WRV5+3Ec5ykHRm9zgULIwquuz5TQ4HVfXUPxRAd5YQH6VRD1YScll3pX3UFSKXokRBlATalbDQVy3pmsomZCkXBiQlcQNUSMWh/iXff0CE+57hQaQpV3XrIPTeG2NRQAZcXACFlAtJte41YeuwIS/ksnVAIuatQxgZ3EmhISY6yfBAioyY7hk+JKzl1FOarnQj8sVa9RnLOzAhhoLMQE3NElYVgkDUgqESRSEBUXHjqE3yzKDMwm5Rc0GsOm2BemPEusrKk7VojvRuqzGP7TI+oSmcv6Jr4gnzs3i9iaw5bY54Y9JKkQ7RHpU7KfWm9/eAlxemPtp4Ge+5kg5cqiSc1ojT8Xrl3/hxqj15i9cwMIn5jEaRJl6nEEzDFBj/mQhMD/IJ+FXHx+fUgjH4H1RVrwWWUU3UfREY36+SKbRHzQHfAoiv0RxpmoecBO14CZfROXSeqGY4gAiKnwZP/QrNg+Oi2RaiDLfw4DKRBnFAOgHSN+U/MiEhMpGGFBYWZ6gRXn9v3AA8vNylYThAB0se05XaohNLJRL5fhd0AX2YHPkl86LDqn19OfLgMst+iENsSHtBaCqOhaghn7LBWzMoUdraAqp7fg2ygXKWLWenrMRuOfQYzWEiYouXoy7B2/j3do4bWo/oKnSuriLcZHRLXp6jjtYn0MPI2hSHpEdL8efRtkRCrLMfCkHvXkclgk5PWc9GGC/q4lMpIcyH1vSRB6RZM5H+8DHWZ6ihGVCvMt+AGBszrHfRYZ3AA/g80i1SA++AF9QZOTqLvuczfD96H7NnvJfPK0z4Hl48KJl8Z+wJS2Wooy0v+YO18RVb5x5ZMKjT8GARlmT+jqjD/CVRXd3YotqRJMHRer8ZhXcW9DmPXb2IKEKqM0rysYpyr9Qd18OK+jHiKZwHlnp/HgeXNYZYPa1nuRyROwp+CLhtLxXOC0vlvdhkREeFHpemw3tCOjTa9/5cAJBWHzryXP5XrkyIiKcwvl9vfpgey0jMKO4pyBCg6ApHs5XYh1qvnzl7msfJS+QpMF1PYtn4FsMhDm+1zdWAhK+KWSyfXjRYeGeAm2svuk1Hn+9w13CuGKrQOKXQAHlSsMFgZ6ega38gVbF49S7ZiZiMB0mEHhqEagLq+LpBi17RSMwoHt3z07SvfsAYuNvBGyvmvl2CrOrtuTGgQFGKcx8fH30xNfHTDGK0rswEeYIhdkPemOmGMHU2iwjDqMVZj/qiZlitIN2y4AAhyEKdu53epDLVgxx0NUwmAxU7Gd37GfvVwzU0XDgLY/t2Kts9m/twmZfHWurs+WBe2YjMHlixml2Qtv2BPbpjImTgZFJ223fJJuM04aGv+vE0PB0hs2ktts+VePZZ5BNxitDyJ8UsMqrDJtBfdptPFURcxe53cn4nI5ZDM9B8IP0zxl33BbN7ciOjSVU0LYO2uF65+8779LTz51LT38HD113DNoKm28jph79O97+I9ttB25zs3HNzna1cds20HY70kb7/z+munmwQ7rqZQAAAABJRU5ErkJggg==",
		Tag:          "DAI",
		Name:         "Dai (DAI)",
		Trezor:       false,
		Ledger:       false,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   true,
		TokenNetwork: "ethereum",
		Contract:     "0x6b175474e89094c44da98b954eedeac495271d0f",
		Decimals:     18,
		Blockbook:    "https://eth2.trezor.io",
	},
	Rates: RatesSource{
		Exchange:         "hitbtc",
		FallBackExchange: "bittrex",
	},
}
