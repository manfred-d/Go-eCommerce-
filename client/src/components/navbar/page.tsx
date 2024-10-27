import { Button } from "../ui/button";
import { Input } from "../ui/input";

const NavBar = () => {
  return (
    <nav className=" w-full flex items-center bg-green-400 justify-between shadow-md">
      <div className="flex flex-col md:flex-row items-center  px-4 py-6  w-full">
        <a href="#" className=" text-3xl text-wrap italic mx-3 pb-5">
          Meet Up
        </a>
        <div className=" px-2 md:w-1/3  w-full">
          <form action="" method="post" className="">
            <div className=" relative">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="25"
                  height="25"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  className="lucide lucide-search absolute left-3 top-1/2 -translate-y-1/2 h-4 w-6 text-muted-foreground text-green-100"
                >
                  <circle cx="11" cy="11" r="8"></circle>
                  <path d="m21 21-4.3-4.3"></path>
                </svg>
                <Input
                  className=" w-full rounded-md bg-transparent px-5 py-5 shadow-sm text-lg transition-colors focus-visible:text-white focus-visible:outline-none pl-10 "
                  type="text"
                  placeholder="search for events ..."
                />
            </div>
          </form>
        </div>
      </div>
      <div className="flex px-6 mx-10">
        <Button variant="secondary">Sign In</Button>
      </div>
    </nav>
  );
}

export default NavBar;