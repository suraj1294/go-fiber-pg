import * as React from "react";

import { Icons } from "@/components/icon";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { SubmitHandler, useForm } from "react-hook-form";
import { useMutation } from "react-query";
import { login as loginUser } from "@/services/auth";
import { AlertDescription } from "@/components/ui/alert";
import { AlertCircle } from "lucide-react";
import { ApiErrorResponse } from "@/models/api-response";
import { Navigate } from "react-router-dom";
import { setAuthState } from "@/auth/auth-store";
import { useAuthStore } from "@/auth";
import { useStore } from "zustand";

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {}

interface AuthForm {
  email: string;
  password: string;
}

export function Login({ className, ...props }: UserAuthFormProps) {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<AuthForm>();

  const store = useAuthStore();
  const token = useStore(store, (state) => state.token);

  const { mutate, isLoading, isError, error } = useMutation(loginUser, {
    onError: (e: ApiErrorResponse) => {
      console.error(e);
    },
    onSuccess: (data) => {
      setAuthState?.({ token: data?.accessToken ?? "" });
    },
  });

  const onSubmit: SubmitHandler<AuthForm> = (data) => {
    const { email, password } = data;
    mutate({ email, password });
  };

  if (token) {
    return <Navigate to="/" />;
  }

  return (
    <section className="max-w-lg mx-auto mt-32">
      <form onSubmit={handleSubmit(onSubmit)}>
        <Card className="w-[350px]">
          <CardHeader>
            <CardTitle>Sign In</CardTitle>
          </CardHeader>
          <CardContent className={className} {...props}>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label className="sr-only" htmlFor="email">
                  Email
                </Label>
                <Input
                  id="email"
                  placeholder="name@example.com"
                  type="email"
                  autoCapitalize="none"
                  autoComplete="email"
                  autoCorrect="off"
                  {...register("email")}
                  disabled={isLoading}
                />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label className="sr-only" htmlFor="password">
                  Password
                </Label>
                <Input
                  id="password"
                  placeholder="password"
                  type="password"
                  autoCapitalize="none"
                  autoComplete="password"
                  autoCorrect="off"
                  {...register("password")}
                  disabled={isLoading}
                />
              </div>
            </div>
            {isError && (
              <div className="flex items-center justify-center pt-2 text-destructive gap-2">
                <AlertCircle className="h-4 w-4" />
                <AlertDescription>{error?.message ?? ""}</AlertDescription>
              </div>
            )}
            {errors.email && (
              <div className="text-destructive">
                <AlertDescription>
                  {errors?.email?.message ?? ""}
                </AlertDescription>
              </div>
            )}
            {errors.password && (
              <div className="text-destructive">
                <AlertDescription>
                  {errors?.password?.message ?? ""}
                </AlertDescription>
              </div>
            )}
          </CardContent>

          <CardFooter className="flex justify-between">
            <Button variant="outline">Cancel</Button>
            <Button disabled={isLoading}>
              {isLoading && (
                <Icons.spinner className="mr-2 h-4 w-4 animate-spin" />
              )}
              Sign In
            </Button>
          </CardFooter>
        </Card>
      </form>
    </section>
  );
}
