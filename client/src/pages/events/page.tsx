import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"

import Image from "../../assets/images/office-business-meeting.jpg";
import { CalendarIcon,  CircleIcon } from "@radix-ui/react-icons";

const events = [
  {
    title: "Droidcon",
    hostedBy: "Android",
  },
  {
    title: "Droidcon",
    hostedBy: "Android",
  },
  {
    title: "Droidcon",
    hostedBy: "Android",
  },
  {
    title: "Droidcon",
    hostedBy: "Android",
  },
  {
    title: "Droidcon",
    hostedBy: "Android",
  }
];

const EventsPage = () => {
  return (
    <div className=" mt-24">
      <div className="headers my-4 mx-6 flex flex-row justify-between items-center">
        <div className="header-title">
          <h3 className=" text-xl text-orange-500 font-bold">
            Upcoming Events
          </h3>
        </div>
        <div className="link">
          <a href="">View all events</a>
        </div>
      </div>
      <div className="events grid grid-flow-row xl:grid-cols-4 md:grid-cols-3 sm:grid-cols-2 gap-4">
        {events.map((event, index) => (
          <Card
            key={index}
            className=" w-[350px] hover:bg-gray-100 hover:shadow-md   "
          >
            <CardHeader>
              <img
                src={Image}
                alt="image about meeting"
                width="400"
                height="300"
                fetchPriority="high"
                decoding="async"
                className=" via-transparent"
              />
            </CardHeader>
            <CardContent className=" grid gap-4">
              <div className="name">
                <p className="text-4xl font-bold text-muted-foreground">
                  {event.title}
                </p>
              </div>
              <div className="hosts flex items-center justify-start flex-row gap-6">
                <p className=" text-md text-gray-500">Hosted By: </p>
                <p className=" text-lg font-medium text-gray-500">
                  {event.hostedBy}{" "}
                </p>
              </div>
              <div className="desc my-2">
                <h5 className=" text-base text-muted-foreground">
                  Learning and development roundtable
                </h5>
              </div>
              <div className="date flex items-center flex-row gap-6">
                <p className=" text-base text-gray-500 flex items-center gap-2">
                  <CalendarIcon /> Mon, Oct 28{" "}
                </p>
                <p className=" text-base text-gray-500">8:00 PM EAT</p>
              </div>
              <div className="attend flex flex-row gap-6">
                <p className=" text-lg text-gray-500 flex items-center gap-2">
                  <CircleIcon /> 10 going{" "}
                </p>
                <p className=" text-lg text-gray-500 flex gap-2 items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-6"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M2.25 18.75a60.07 60.07 0 0 1 15.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 0 1 3 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 0 0-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 0 1-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 0 0 3 15h-.75M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Zm3 0h.008v.008H18V10.5Zm-12 0h.008v.008H6V10.5Z"
                    />
                  </svg>
                  From $5.00
                </p>
              </div>
            </CardContent>
            <CardFooter>
              <Button className="w-full" variant="outline">
                View Event Details
              </Button>
            </CardFooter>
          </Card>
        ))}
      </div>
    </div>
  );
}

export default EventsPage
