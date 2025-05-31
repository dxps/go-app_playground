package comps

const (
	PageCss                = "flex flex-col min-h-screen bg-gray-100"                                 // the page's container (first `div` element)
	MainContentCss         = "flex flex-col min-h-screen justify-center items-center drop-shadow-2xl" // main content (included after Navbar)
	MainContentChildDivCss = "bg-white rounded-md p-6 min-w-[600px]"                                  // the child (a `div` element) of the main content container

	TitleCss            = "flex-grow text-lg font-medium px-2 text-center text-gray-400"                                                                         // the title (shown on various pages)
	CloseBtnCss         = "bg-gray-100 hover:bg-green-100 disabled:text-gray-400 disabled:hover:bg-gray-200 drop-shadow-sm px-4 py-1 rounded-md"                 // the (button-right) close button
	CloseSymbolCss      = "text-gray-400 text-xl hover:text-gray-900 hover:bg-gray-100 px-2 rounded-xl transition duration-200 cursor-pointer"                   // the (top-right) close symbol
	PositiveBtnCss      = "bg-gray-100 bg-green-100 enabled:hover:bg-green-100 disabled:text-gray-400 hover:disabled:bg-gray-100 drop-shadow-sm px-4 rounded-md" // a "positive" (due to its green bg)
	DeleteBtnCss        = "text-red-200 bg-slate-50 hover:text-red-600 hover:bg-red-100 drop-shadow-sm px-4 rounded-md"                                          // the (bottom-left) delete button
	DeleteBtnConfirmCss = "text-red-500 bg-red-100 hover:text-red-800 hover:bg-red-200 drop-shadow-sm px-4 rounded-md"                                           // the (bottom-left) delete button

	FormLabelCss        = "pr-3 py-2 min-w-40 text-gray-600"
	FormFieldCss        = "px-3 py-2 my-1 rounded-lg outline-none border-1 focus:border-green-300 min-w-80"
	FormBooleanFieldCss = "px-3 py-2 my-1 rounded-lg outline-none border-1 focus:border-green-300"
)
