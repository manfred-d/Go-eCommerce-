import { Button } from "@/components/ui/button";
import Image from '../../assets/images/office-business-meeting.jpg'
import EventsPage from "../events/page";
import TopCategories from "@/components/categories/page";
const HomePage = () => {
  return (
    <div className=" relative m-w-screen flex flex-col flex-grow items-center justify-center pt-20 sm:px-10 sm:pb-8">
        <div className=" mb-10 sm:w-[580px] md:w-[760px] xl:w-[1200px] gap-10  px-4 w-full h-full min-h-full flex flex-col lg:flex-row items-center justify-center">
          <div className="content row w-full sm:w-full lg:w-7/12 md:pr-10">
            <div className="h">
              <h2 className=" md:text-[42px] text-2xl font-black  pb-4 lg:leading-[3.25rem]">
                The people platform - Where people meet and share ideas, events and
                activities. <br />
                <span className=" text-2xl">Making a long lasting friendships.</span>
              </h2>
              <p className="text-lg font-normal mb-10">
                Whatever your interest, from hiking and reading to networking and
                skill sharing, there are thousands of people who share it on Meetup.
                Events are happening every day—sign up to join the fun.
              </p>
              <Button variant="ghost" className=" py-6 w-full text-white text-xl">Join Meet Up</Button>
            </div>
          </div>
          <div className="content w-full sm:w-full lg:w-1/2 md:pl-10 ">
             <img src={Image} alt="image about meeting" width="500" height="400" fetchPriority="high" decoding="async" className=" via-transparent"  />
          </div>
        </div>
        <EventsPage />
        <TopCategories />
    </div>
  );
}

export default HomePage
