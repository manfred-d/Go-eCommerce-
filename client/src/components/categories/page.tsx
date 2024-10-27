import { PieChartIcon, SunIcon } from "@radix-ui/react-icons";
import { Card, CardContent, CardHeader } from "../ui/card";


const categories = [
  {
    title: "Social Activities",
  },
  {
    title: "Technology",
  },
  {
    title: "Sports and Fitness",
  },
  {
    title: "Arts and Culture",
  },
  {
    title: "Health Events",
  },
  {
    title: "Games",
  },
  {
    title: "Travel and Leisure",
  },
  {
    title: "Hobbies and Passions",
  },
];

function TopCategories() {
  return (
    <div className=" my-10">
      <div className="headers my-4 mx-6 flex flex-row justify-between items-center">
        <div className="header-title">
          <h3 className=" text-xl text-orange-500 font-bold">
            Upcoming Events
          </h3>
        </div>
      </div>
      <div className="cat flex gap-4 ">
        {categories.map((categori, index) => (
          <Card
            key={index}
            className=" w-[180px] h-40 flex flex-col items-center text-center gap-4 hover:bg-gray-100 hover:shadow-md shadow outline-none p-2 "
          >
            <CardHeader className=" py-2 px-0 ">
              <SunIcon className=" w-14 h-10 " />
            </CardHeader>
            <CardContent>
              <span className=" text-lg text-muted-foreground font-medium ">
                {categori.title}
              </span>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}

export default TopCategories;
