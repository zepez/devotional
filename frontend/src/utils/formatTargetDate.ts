import { format } from "date-fns";


const formatTargetDate = (targetDate:string) => {
	if(!targetDate) return "Unknown";
	if(targetDate.length < 8) return "Unknown";

	const year = parseInt(targetDate.substring(0, 4));
	const month = parseInt(targetDate.substring(4, 6)) - 1;
	const day = parseInt(targetDate.substring(6, 8));
	console.log(year, month, day);

	const formatted = format(new Date(year, month, day), "EEEE, MMMM do yyyy");

	console.log(formatted);
	return formatted;
};


export default formatTargetDate;
