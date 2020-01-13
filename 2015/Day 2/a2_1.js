function compareNumbers(a, b) {
  return a - b;
}
var fs = require('fs');
var rawFile;
fs.readFile('a2.txt','utf-8',function(err,data){
	rawFile = data;
	var total = 0;
	rawFile = rawFile.split('\r\n');
	for(var i=0;i<rawFile.length;i++){
		rawFile[i]=rawFile[i].split('x');
		var temp=[];
		for(var j=0;j<3;j++){
			temp.push(Number(rawFile[i][j]));
		}
		temp = temp.sort(compareNumbers);
		var temp_total = (3*(temp[0])*(temp[1])) + (2*(temp[1])*(temp[2])) + (2*(temp[2])*(temp[0]))
		total = total + temp_total;
	}
	console.log(total);
});