var crypto = require('crypto');
var gi = "ckczppom";
var found = false;
for(var i=0; found!= true ; i++){
	var md5sum = crypto.createHash('md5');
	var gi2 = gi + i;
	md5sum.update(gi2);
	var se = md5sum.digest('hex');
	if(se.slice(0,5)==="00000"){
		found = true;
		console.log(i);
	}
}