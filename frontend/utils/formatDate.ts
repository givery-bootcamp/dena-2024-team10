import dayjs from "dayjs";

const formatDate = (date: string) => {
	return dayjs(date).format("YYYY/MM/DD HH:mm");
};

export default formatDate;
