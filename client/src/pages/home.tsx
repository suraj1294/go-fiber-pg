import { UserNav } from "@/components/header/user-menu";
import { MainNav } from "@/components/main-nav";
import { Search } from "@/components/search";
import { Button } from "@/components/ui/button";
import { getMovies } from "@/services/movies";
import { FC } from "react";
import { useQuery } from "react-query";

const Home: FC = () => {
  const response = useQuery("getMovies", getMovies);

  console.log(response.data?.data.response);

  return (
    <div className="border-b">
      <div className="flex h-16 items-center px-4">
        <MainNav className="mx-6" />
        <div className="ml-auto flex items-center space-x-4">
          <Search />
          <UserNav />
          <Button onClick={() => response.refetch()}>Refetch</Button>
        </div>
      </div>
    </div>
  );
};

export default Home;
