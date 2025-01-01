
import SuperTokens, { SuperTokensWrapper } from "supertokens-auth-react";
import ThirdParty, { Github, Google, Facebook, Apple } from "supertokens-auth-react/recipe/thirdparty";
import EmailPassword from "supertokens-auth-react/recipe/emailpassword";
import Session, { SessionAuth } from "supertokens-auth-react/recipe/session";
import { PrimeReactProvider } from "primereact/api";
import { BrowserRouter, Route, Routes } from "react-router";
import * as reactRouterDom from "react-router";
import { getSuperTokensRoutesForReactRouterDom } from "supertokens-auth-react/ui";
import { ThirdPartyPreBuiltUI } from "supertokens-auth-react/recipe/thirdparty/prebuiltui";
import { EmailPasswordPreBuiltUI } from "supertokens-auth-react/recipe/emailpassword/prebuiltui";
import { Home } from "./pages/Home";


SuperTokens.init({
  appInfo: {
    // learn more about this on https://supertokens.com/docs/thirdpartyemailpassword/appinfo
    appName: "My Server",
    apiDomain: "http://localhost:8080",
    websiteDomain: "http://localhost:5173",
    apiBasePath: "/api/auth",
    websiteBasePath: "/auth"
  },
  recipeList: [
    ThirdParty.init({
      signInAndUpFeature: {
        providers: [
          Github.init(),
          Google.init(),
          Facebook.init(),
          Apple.init(),
        ]
      }
    }),
    EmailPassword.init(),
    Session.init()
  ]
});


function App() {

  return (
    <>
      <SuperTokensWrapper>
        <PrimeReactProvider>
          <BrowserRouter>
            <Routes>
              {/*This renders the login UI on the /auth route*/}
              {getSuperTokensRoutesForReactRouterDom(reactRouterDom, [ThirdPartyPreBuiltUI, EmailPasswordPreBuiltUI])}
              {/*App routes*/}

              <Route index element={
                <SessionAuth>
                  <Home />
                </SessionAuth>
              }
              />
            </Routes>
          </BrowserRouter>
        </PrimeReactProvider>

      </SuperTokensWrapper>
    </>
  )
}

export default App
