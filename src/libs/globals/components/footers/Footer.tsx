import { joinClass } from "@/libs/globals/utilities/joinClass";
import { FC, Fragment } from "react";
import { Facebook, Instagram, Youtube } from "react-feather";
import FooterIcon from "./FooterIcon";
import FooterTitle from "./FooterTitle";
import CopyrightFooter from "./CopyrightFooter";
import Button from "../buttons/Button";
import FooterColumn from "./FooterColumn";
import FooterLink from "./FooterLink";

const Footer: FC = () => {
  return (
    <Fragment>
      <div
        className={joinClass(
          "flex w-full flex-col gap-6",
          "bg-(--bg-main) px-12 py-8 md:flex-row",
        )}
      >
        <div className="flex flex-1 flex-col gap-2 md:mr-36">
          <FooterTitle label="Follow Us" />
          <div className="flex flex-row gap-4 text-black">
            <FooterIcon Icon={Facebook} />
            <FooterIcon Icon={Instagram} />
            <FooterIcon Icon={Youtube} />
          </div>
          <div className="mt-16 flex flex-col gap-2">
            <FooterTitle label="Get access to exclusive offers and deals" />
            <form className="flex w-full">
              <input
                className="flex-1 rounded-xs bg-white p-3 placeholder:text-slate-400"
                type="email"
                placeholder="E-mail"
                name="userEmail"
                id="userEmail"
              />
              <Button type="submit" className="px-6 py-2" variant="contained">
                Subscribe
              </Button>
            </form>
          </div>
        </div>
        <div className="hidden flex-2 flex-row gap-6 lg:flex">
          <FooterColumn label="About us">
            <FooterLink label="Home" />
            <FooterLink label="Contact" />
            <FooterLink label="Sustainability" />
            <FooterLink label="Press" />
          </FooterColumn>
          <FooterColumn label="Help and support">
            <FooterLink label="Deliver" />
            <FooterLink label="Payments" />
            <FooterLink label="Contact" />
            <FooterLink label="Care Instructions" />
            <FooterLink label="FAQ" />
          </FooterColumn>
          <FooterColumn label="Term of use">
            <FooterLink label="Warranty" />
            <FooterLink label="Return Policy" />
            <FooterLink label="Details" />
          </FooterColumn>
          <FooterColumn label="Why HLI LIGHT?">
            <FooterLink label="Free Shipping" />
            <FooterLink label="Play by invoice" />
            <FooterLink label="Return Within 365 Days" />
            <FooterLink label="Up to 70% Discounts" />
            <FooterLink label="Warranty up to 5 years" />
          </FooterColumn>
        </div>
      </div>
      <CopyrightFooter />
    </Fragment>
  );
};

export default Footer;
