var fs = require('fs');
var rawFile;
fs.readFile('a3.txt','utf-8',function(err,data){
	rawFile = data;
	var total = 0;
	var cx=0;
	var cy=0;
	var temp = new Set;
	temp.add(cx+"+"+cy);
	for(var i=0;i<rawFile.length;i+=1){
		if(rawFile[i]==='^'){cy++;}
		else if(rawFile[i]==='<'){cx--;}
		else if(rawFile[i]==='>'){cx++;}
		else if(rawFile[i]==='v'){cy--;}
		temp.add(cx+"+"+cy);
	}
	total = temp.size
	console.log(total);
});