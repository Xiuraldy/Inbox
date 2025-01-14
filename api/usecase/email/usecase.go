package usecase

import repository "challenge/repository/email"

type UseCase interface {
	GetEmail() (map[string]interface{}, error)
}

type useCaseEmail struct {
	RepositoryEmail repository.Repository
}

func NewUseCaseEmail(repositoryEmail repository.Repository) UseCase {
	return &useCaseEmail{
		RepositoryEmail: repositoryEmail,
	}
}

func (u *useCaseEmail) GetEmail() (map[string]interface{}, error) {
	data := []interface{}{
		map[string]interface{}{
			"_timestamp": 1736025452012954,
			"body":       "---------------------- Forwarded by Peter Parker/HOU/ECT on 09/17/2000 10:00 AM ---------------------------\n\n\n\"Tony Stark\" <tony.stark@provider.com> on 09/15/2000 06:45:00 PM\nPlease respond to <tony.stark@provider.com>\nTo: \"Peter Parker\" <pparker@enron.com>\ncc: \"Steve Rogers\" <srogers@mail.sanmarcos.net> \nSubject: Stark Tower Proposal - Peter Parker.pdf\n\nThe attached proposal outlines the redevelopment of Stark Tower. Anticipated revenue growth of 40% over the next 3 years.",
			"date":       "Sat, 17 Sep 2000 10:00:00 -0700 (PDT)",
			"from":       "peter.parker@enron.com",
			"message_id": "<1357912.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Stark Tower Proposal - Peter Parker.pdf",
			"to":         "steve.rogers@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025462023955,
			"body":       "---------------------- Forwarded by Diana Prince/HOU/ECT on 09/18/2000 11:45 AM ---------------------------\n\n\n\"Arthur Curry\" <arthur.curry@provider.com> on 09/16/2000 08:15:00 PM\nPlease respond to <arthur.curry@provider.com>\nTo: \"Diana Prince\" <dprince@enron.com>\ncc: \"Barry Allen\" <ballen@mail.sanmarcos.net> \nSubject: Atlantis Expansion - Diana Prince.xls\n\nAttached is the strategic plan for the Atlantis expansion. ROI projected at 35% over 4 years.",
			"date":       "Sun, 18 Sep 2000 11:45:00 -0700 (PDT)",
			"from":       "diana.prince@enron.com",
			"message_id": "<2468102.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Atlantis Expansion - Diana Prince.xls",
			"to":         "barry.allen@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025472034956,
			"body":       "---------------------- Forwarded by Natasha Romanoff/HOU/ECT on 09/19/2000 02:30 PM ---------------------------\n\n\n\"Clint Barton\" <clint.barton@provider.com> on 09/17/2000 07:50:00 AM\nPlease respond to <clint.barton@provider.com>\nTo: \"Natasha Romanoff\" <nromanoff@enron.com>\ncc: \"Bruce Banner\" <bbanner@mail.sanmarcos.net> \nSubject: Avengers Initiative - Natasha Romanoff.doc\n\nDraft of the Avengers Initiative plan is attached. The project aims for a 60% success rate within 5 years.",
			"date":       "Mon, 19 Sep 2000 14:30:00 -0700 (PDT)",
			"from":       "natasha.romanoff@enron.com",
			"message_id": "<3579132.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Avengers Initiative - Natasha Romanoff.doc",
			"to":         "bruce.banner@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025482045957,
			"body":       "---------------------- Forwarded by Logan Howlett/HOU/ECT on 09/20/2000 09:15 AM ---------------------------\n\n\n\"Charles Xavier\" <charles.xavier@provider.com> on 09/18/2000 05:30:00 PM\nPlease respond to <charles.xavier@provider.com>\nTo: \"Logan Howlett\" <lhowlett@enron.com>\ncc: \"Erik Lehnsherr\" <elehnsherr@mail.sanmarcos.net> \nSubject: Xavier's School Expansion - Logan Howlett.xlsx\n\nProposal for the expansion of Xavier's School. Anticipated outcomes include 50% student growth over 2 years.",
			"date":       "Tue, 20 Sep 2000 09:15:00 -0700 (PDT)",
			"from":       "logan.howlett@enron.com",
			"message_id": "<4680243.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Xavier's School Expansion - Logan Howlett.xlsx",
			"to":         "erik.lehnsherr@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025492056958,
			"body":       "---------------------- Forwarded by Hermione Granger/HOU/ECT on 09/22/2000 11:45 AM ---------------------------\n\n\n\"Harry Potter\" <harry.potter@provider.com> on 09/20/2000 08:15:00 PM\nPlease respond to <harry.potter@provider.com>\nTo: \"Hermione Granger\" <hgranger@enron.com>\ncc: \"Ron Weasley\" <rweasley@mail.sanmarcos.net> \nSubject: Hogwarts Expansion - Hermione Granger.xls\n\nAttached is the strategic plan for Hogwarts' expansion. Projected ROI is 25% annually over 3 years.",
			"date":       "Fri, 22 Sep 2000 11:45:00 -0700 (PDT)",
			"from":       "hermione.granger@enron.com",
			"message_id": "<2468103.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Hogwarts Expansion - Hermione Granger.xls",
			"to":         "ron.weasley@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025502067959,
			"body":       "---------------------- Forwarded by Bruce Banner/HOU/ECT on 09/23/2000 01:30 PM ---------------------------\n\n\n\"Steve Rogers\" <steve.rogers@provider.com> on 09/21/2000 06:45:00 PM\nPlease respond to <steve.rogers@provider.com>\nTo: \"Bruce Banner\" <bbanner@enron.com>\ncc: \"Tony Stark\" <tstark@mail.sanmarcos.net> \nSubject: Shield Proposal - Bruce Banner.pdf\n\nEnclosed is the Shield project proposal. ROI estimates suggest a return of 45% over 3 years.",
			"date":       "Sat, 23 Sep 2000 13:30:00 -0700 (PDT)",
			"from":       "bruce.banner@enron.com",
			"message_id": "<5791314.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Shield Proposal - Bruce Banner.pdf",
			"to":         "tony.stark@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025512078960,
			"body":       "---------------------- Forwarded by Wanda Maximoff/HOU/ECT on 09/24/2000 04:00 PM ---------------------------\n\n\n\"Vision\" <vision@provider.com> on 09/22/2000 07:15:00 PM\nPlease respond to <vision@provider.com>\nTo: \"Wanda Maximoff\" <wmaximoff@enron.com>\ncc: \"Pietro Maximoff\" <pmaximoff@mail.sanmarcos.net> \nSubject: Hex Expansion - Wanda Maximoff.doc\n\nProposal for Hex expansion is attached. Anticipated ROI is 30% over the next 4 years.",
			"date":       "Sun, 24 Sep 2000 16:00:00 -0700 (PDT)",
			"from":       "wanda.maximoff@enron.com",
			"message_id": "<6802425.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Hex Expansion - Wanda Maximoff.doc",
			"to":         "pietro.maximoff@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025522089961,
			"body":       "---------------------- Forwarded by Sam Wilson/HOU/ECT on 09/25/2000 11:45 AM ---------------------------\n\n\n\"Bucky Barnes\" <bucky.barnes@provider.com> on 09/23/2000 08:15:00 AM\nPlease respond to <bucky.barnes@provider.com>\nTo: \"Sam Wilson\" <swilson@enron.com>\ncc: \"Steve Rogers\" <srogers@mail.sanmarcos.net> \nSubject: Winter Soldier Project - Sam Wilson.pdf\n\nProposal for Winter Soldier Project. Expected annual growth is projected at 28% for the next 5 years.",
			"date":       "Mon, 25 Sep 2000 11:45:00 -0700 (PDT)",
			"from":       "sam.wilson@enron.com",
			"message_id": "<7813536.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Winter Soldier Project - Sam Wilson.pdf",
			"to":         "steve.rogers@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025532100962,
			"body":       "---------------------- Forwarded by Carol Danvers/HOU/ECT on 09/26/2000 09:15 AM ---------------------------\n\n\n\"Nick Fury\" <nick.fury@provider.com> on 09/24/2000 06:30:00 PM\nPlease respond to <nick.fury@provider.com>\nTo: \"Carol Danvers\" <cdanvers@enron.com>\ncc: \"Maria Rambeau\" <mrambeau@mail.sanmarcos.net> \nSubject: Captain Marvel Initiative - Carol Danvers.doc\n\nEnclosed is the detailed Captain Marvel Initiative. Anticipated success rate of 75% within 3 years.",
			"date":       "Tue, 26 Sep 2000 09:15:00 -0700 (PDT)",
			"from":       "carol.danvers@enron.com",
			"message_id": "<8914647.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Captain Marvel Initiative - Carol Danvers.doc",
			"to":         "maria.rambeau@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025542111963,
			"body":       "---------------------- Forwarded by Peter Quill/HOU/ECT on 09/27/2000 01:30 PM ---------------------------\n\n\n\"Rocket Raccoon\" <rocket.raccoon@provider.com> on 09/25/2000 08:00:00 AM\nPlease respond to <rocket.raccoon@provider.com>\nTo: \"Peter Quill\" <pquill@enron.com>\ncc: \"Gamora\" <gamora@mail.sanmarcos.net> \nSubject: Guardians Initiative - Peter Quill.xls\n\nAttached is the Guardians Initiative strategic plan. Expected annual growth is 35% for the next 5 years.",
			"date":       "Wed, 27 Sep 2000 13:30:00 -0700 (PDT)",
			"from":       "peter.quill@enron.com",
			"message_id": "<9015758.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Guardians Initiative - Peter Quill.xls",
			"to":         "gamora@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025552122964,
			"body":       "---------------------- Forwarded by T'Challa/HOU/ECT on 09/28/2000 02:45 PM ---------------------------\n\n\n\"Shuri\" <shuri@provider.com> on 09/26/2000 07:00:00 AM\nPlease respond to <shuri@provider.com>\nTo: \"T'Challa\" <tchalla@enron.com>\ncc: \"Okoye\" <okoye@mail.sanmarcos.net> \nSubject: Wakanda Expansion Plan - T'Challa.doc\n\nDetails of Wakanda's Expansion Plan attached. Expected revenue growth of 50% over the next 4 years.",
			"date":       "Thu, 28 Sep 2000 14:45:00 -0700 (PDT)",
			"from":       "tchalla@enron.com",
			"message_id": "<1126979.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Wakanda Expansion Plan - T'Challa.doc",
			"to":         "okoye@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025401967949,
			"body":       "---------------------- Forwarded by John Doe/HOU/ECT on 09/12/2000 04:57 PM ---------------------------\n\n\n\"Jane Smith\" <jane.smith@provider.com> on 09/10/2000 06:21:49 AM\nPlease respond to <jane.smith@provider.com>\nTo: \"John Doe\" <jdoe@enron.com>\ncc: \"Alan Brown\" <abrown@mail.sanmarcos.net> \nSubject: Project Analysis - John Doe.xls\n\nEnclosed is the preliminary analysis for the City Center project in Dallas. As you can see, this project shows a potential return of over 30% annually for 4 years.",
			"date":       "Mon, 12 Sep 2000 10:57:00 -0700 (PDT)",
			"from":       "john.doe@enron.com",
			"message_id": "<1234567.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Project Analysis - John Doe.xls",
			"to":         "alan.brown@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025411978950,
			"body":       "---------------------- Forwarded by Sarah Connor/HOU/ECT on 09/13/2000 11:15 AM ---------------------------\n\n\n\"Tom Johnson\" <tom.johnson@provider.com> on 09/11/2000 07:30:25 PM\nPlease respond to <tom.johnson@provider.com>\nTo: \"Sarah Connor\" <sconnor@enron.com>\ncc: \"Ellen White\" <ewhite@mail.sanmarcos.net> \nSubject: Investment Proposal - Sarah Connor.pdf\n\nAttached is the investment proposal for the Oakwood development in Houston, promising a steady 20% ROI over 5 years.",
			"date":       "Tue, 13 Sep 2000 11:15:00 -0700 (PDT)",
			"from":       "sarah.connor@enron.com",
			"message_id": "<7654321.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Investment Proposal - Sarah Connor.pdf",
			"to":         "ellen.white@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025421989951,
			"body":       "---------------------- Forwarded by Michael Scott/HOU/ECT on 09/14/2000 03:45 PM ---------------------------\n\n\n\"Dwight Schrute\" <dwight.schrute@provider.com> on 09/12/2000 05:45:00 PM\nPlease respond to <dwight.schrute@provider.com>\nTo: \"Michael Scott\" <mscott@enron.com>\ncc: \"Jim Halpert\" <jhalpert@mail.sanmarcos.net> \nSubject: Expansion Plan - Michael Scott.xlsx\n\nDetails of the expansion plan for the Scranton project have been included. Anticipated ROI exceeds 25% over 3 years.",
			"date":       "Wed, 14 Sep 2000 15:45:00 -0700 (PDT)",
			"from":       "michael.scott@enron.com",
			"message_id": "<9876543.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Expansion Plan - Michael Scott.xlsx",
			"to":         "jim.halpert@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025431990952,
			"body":       "---------------------- Forwarded by Laura Palmer/HOU/ECT on 09/15/2000 01:30 PM ---------------------------\n\n\n\"Benjamin Horne\" <ben.horne@provider.com> on 09/13/2000 08:00:00 PM\nPlease respond to <ben.horne@provider.com>\nTo: \"Laura Palmer\" <lpalmer@enron.com>\ncc: \"Donna Hayward\" <dhayward@mail.sanmarcos.net> \nSubject: Twin Peaks Proposal - Laura Palmer.doc\n\nAttached is the development proposal for the Twin Peaks project. Expected revenue growth is projected at 15% annually for the next 6 years.",
			"date":       "Thu, 15 Sep 2000 13:30:00 -0700 (PDT)",
			"from":       "laura.palmer@enron.com",
			"message_id": "<2468101.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Twin Peaks Proposal - Laura Palmer.doc",
			"to":         "donna.hayward@enron.com",
		},
		map[string]interface{}{
			"_timestamp": 1736025442001953,
			"body":       "---------------------- Forwarded by Clark Kent/HOU/ECT on 09/16/2000 04:00 PM ---------------------------\n\n\n\"Bruce Wayne\" <bruce.wayne@provider.com> on 09/14/2000 07:10:00 AM\nPlease respond to <bruce.wayne@provider.com>\nTo: \"Clark Kent\" <ckent@enron.com>\ncc: \"Lois Lane\" <llane@mail.sanmarcos.net> \nSubject: Gotham Redevelopment - Clark Kent.pdf\n\nPlease find the attached redevelopment plan for Gotham City. The project forecasts a 50% increase in property value over 2 years.",
			"date":       "Fri, 16 Sep 2000 16:00:00 -0700 (PDT)",
			"from":       "clark.kent@enron.com",
			"message_id": "<1357911.1075855688335.JavaMail.evans@thyme>",
			"subject":    "Gotham Redevelopment - Clark Kent.pdf",
			"to":         "lois.lane@enron.com",
		},
	}

	return map[string]interface{}{
		"data":       data,
		"totalCount": len(data),
	}, nil
}
