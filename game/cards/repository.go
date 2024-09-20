package cards

import (
	"duel-masters/game/cards/dm01"
	"duel-masters/game/cards/dm02"
	"duel-masters/game/cards/dm03"
	"duel-masters/game/cards/dm04"
	"duel-masters/game/cards/dm05"
	"duel-masters/game/cards/dm06"
	"duel-masters/game/cards/dm07"
	"duel-masters/game/match"
)

// Sets is a map of pointers to the available card sets
var Sets = map[string]*map[string]match.CardConstructor{
	"dm-01": &DM01,
	"dm-02": &DM02,
	"dm-03": &DM03,
	"dm-04": &DM04,
	"dm-05": &DM05,
	"dm-06": &DM06,
	"dm-07": &DM07,
}

// DM01 is a map with all the card id's in the game and corresponding CardConstructor for dm01
var DM01 = map[string]match.CardConstructor{

	"57eeb3c3-2561-4841-a381-2e50d17533d1": dm01.AquaHulcus,
	"ecd1ae69-4f63-4e8d-a3f4-9a5c81f98a20": dm01.EmeraldGrass,
	"09b218fc-9c5a-48ef-9555-4908932271e9": dm01.AquaKnight,
	"4097a036-a775-4218-9a1d-f57ead85dda6": dm01.AquaSniper,
	"c43bc627-9e7a-4686-9d61-789425669b02": dm01.AquaSoldier,
	"9781089f-1aa9-4a75-b106-35e9d431e31d": dm01.AquaVehicle,
	"1d72eb3e-5185-449a-a16f-391bd2338343": dm01.BurningMane,
	"fcd0cb50-b687-4180-90a8-390aeb8705cc": dm01.FearFang,
	"10e0e90f-ad7d-4b69-98d5-f01525eb1cdd": dm01.SteelSmasher,
	"015fd6bb-37a9-45cf-bb6b-a5497412b880": dm01.BronzeArmTribe,
	"6663848d-035e-44b6-9d9f-7b236ea5bc43": dm01.GoldenWingStriker,
	"0e26fe1a-a9d1-4c78-80e9-7f4cc0e4c1c8": dm01.MightyShouter,
	"0b1e4f56-6342-46db-9faf-882fd1f1f179": dm01.ArtisanPicora,
	"983e72d7-3f4e-466d-a4e3-06552e392af2": dm01.NomadHeroGigio,
	"0cc5279e-0a26-41a8-a2a5-f7711120b772": dm01.LahPurificationEnforcer,
	"808ddd60-e8ca-49f0-9baa-57e632f85b28": dm01.RaylaTruthEnforcer,
	"91db2302-6794-4aa4-b17b-6637d356e9ac": dm01.AstrocometDragon,
	"0ffdcae3-9db2-401b-8a82-dfad707b83cd": dm01.BolshackDragon,
	"6cf85053-abaa-4577-b151-86123004980e": dm01.Draglide,
	"3b6e6c29-017d-41b9-bf93-186f7963723e": dm01.GatlingSkyterror,
	"1c5511be-7629-41c5-bf17-4bc810be5472": dm01.ScarletSkyterror,
	"a4adb373-0aec-4fff-997c-3820c7ec528d": dm01.DomeShell,
	"1ecb54a2-bcbf-4396-bf09-50dfe984e287": dm01.StormShell,
	"c761c174-87c3-4f4a-ab94-aa837c5ab587": dm01.TowerShell,
	"18e0e199-7827-4a4c-a37d-3acfa4e500d6": dm01.RoaringGreatHorn,
	"2aeae452-5630-4f86-b073-7e9dc07adc43": dm01.StampedingLonghorn,
	"84e1b416-c2d5-4ae1-aca0-025651c6aa58": dm01.TrihornShepherd,
	"3e2940f4-5654-4456-bfc2-fa5e43911cfb": dm01.KingCoral,
	"cd13f7c2-aa5e-43b8-8811-700f230a5de5": dm01.KingDepthcon,
	"f04feb7f-971f-4192-893a-46c23180233a": dm01.KingRippedHide,
	"596f5b72-2502-4120-81f9-9ff9a17271d8": dm01.CandyDrop,
	"a3cf18f0-b04f-45e9-97f7-2a2ead0a1787": dm01.FaerieChild,
	"3f331274-f5f8-42e7-9f28-ce637add34d4": dm01.MarineFlower,
	"ce48ff2c-ea9e-4c12-8629-028d2480b063": dm01.IllusionaryMerfolk,
	"4b021e6f-39cf-401e-89cf-f164f7c0a797": dm01.PhantomFish,
	"cfe9f5b8-2eeb-42c9-89ff-7e69734adc4d": dm01.RevolverFish,
	"70e6cc2c-c63d-4dd9-9b6e-0713fed174bb": dm01.SaucerHeadShark,
	"4c9acf76-cc52-44c3-9e39-613d744c63c5": dm01.PoisonousMushroom,
	"cc9762c3-515a-4734-a3fe-1e0c4c3b3d71": dm01.BoneAssassin,
	"4d3201e8-0d9b-481e-b8e3-86cb90058e20": dm01.BoneSpider,
	"ec46daa1-49ce-4b88-b2bc-e923672ad0f3": dm01.SkeletonSoldierTheDefiled,
	"90b2ed59-828c-4237-ac2e-b7008a02ad2e": dm01.WanderingBraineater,
	"5d3d7052-e5fa-4502-8d31-c72673232317": dm01.HanusaRadianceElemental,
	"6a4270cf-f3be-4c66-8b30-eb2c769065dc": dm01.IocantTheOracle,
	"25a2af16-cc42-4f4c-8c3d-59fb3a7ca74b": dm01.UrthPurifyingElemental,
	"c4839847-e393-47b0-b172-95531aa6d39e": dm01.Gigaberos,
	"5d73062e-acff-47e6-b49a-c0bb1a1762b5": dm01.Gigagiele,
	"6161e271-5294-4073-94d2-b9c06f9d8fa3": dm01.Gigargon,
	"dc1b51b3-52e7-4f1c-8770-515d4e1cb53d": dm01.DeathligerLionOfChaos,
	"07a0115e-797a-49d8-90bf-9ea6de39978d": dm01.ZagaanKnightOfDarkness,
	"7a6f1c82-a8ac-4646-b3e9-fb8592bdd0a4": dm01.Tropico,
	"b7d11c62-2ab3-439b-b147-ae29d34e9216": dm01.FreiVizierOfAir,
	"578ed21b-8ba5-42b2-b662-87a321ee0c7d": dm01.IereVizierOfBullets,
	"cf5eb3d3-e128-42db-bf1a-161d5dd4b972": dm01.LokVizierOfHunting,
	"15efe8b0-02c1-439b-8e7c-4548e74f5c33": dm01.MieleVizierOfLightning,
	"340ec79b-3a4e-4483-ac0e-5dd6b40eb4e1": dm01.ToelVizierOfHope,
	"d067285f-10ea-4666-99c8-bc23e27e3262": dm01.Meteosaur,
	"41c0664d-1969-487d-bde5-866127c1c49e": dm01.Stonesaur,
	"c1ebdda0-be88-4665-937e-2ef3ada8d378": dm01.DeathbladeBeetle,
	"43abeec5-0597-43b3-93cf-766b95d19b5b": dm01.ForestHornet,
	"b3ca1944-41a2-4939-ae85-1a73b1fe085f": dm01.RedEyeScorpion,
	"162f70fb-33f7-4436-a114-41f255c0ce7e": dm01.DarkRavenShadowOfGrief,
	"ea878730-fde0-4bd0-ad25-95e49f54a1b2": dm01.MaskedHorrorShadowOfScorn,
	"f16795cc-4378-4e36-b13a-19f9b932228c": dm01.NightMasterShadowOfDecay,
	"c5a869f4-a959-4667-a352-92df5369e0b9": dm01.DeadlyFighterBraidClaw,
	"f682051b-7cc3-4155-aa8b-eb3335b0435c": dm01.ExplosiveFighterUcarn,
	"c782edd9-34ef-47f5-8f16-af2c3b107a36": dm01.FireSweeperBurningHellion,
	"198ffce7-3d79-420e-9d9b-ebd6421adb6f": dm01.OnslaughterTriceps,
	"becd0856-fb8b-46fd-a950-b57cc5d17c70": dm01.SuperExplosiveVolcanodon,
	"616c146e-049f-4720-a225-0a189729ca79": dm01.ChiliasTheOracle,
	"7b58e8c2-0b1e-4ef5-812f-e667c2092c73": dm01.ReusolTheOracle,
	"d5d57060-ca58-48e1-8903-9b8362c92b0d": dm01.RubyGrass,
	"725a28b7-8c06-4691-93d8-1c6b0dacdba5": dm01.SenatineJadeTree,
	"ae66061e-6039-4dee-abf0-51169913bb35": dm01.ArmoredWalkerUrherion,
	"5370bad9-1260-455e-8120-ea89badc7eaf": dm01.BrawlerZyler,
	"ebd730e1-1099-41ec-a028-6ef1d4cf91b2": dm01.FatalAttackerHorvath,
	"af3bc221-1cc2-4f58-83ea-2673ac2c66c5": dm01.ImmortalBaronVorg,
	"f7dc24d2-2a84-46ff-9661-0b8418d68650": dm01.DiaNorkMoonlightGuardian,
	"39090f65-779c-46c9-856c-67303dd5605c": dm01.GranGureSpaceGuardian,
	"c05fe45d-690e-4856-bddb-5f46154e57e5": dm01.LaUraGigaSkyGuardian,
	"eccceb7c-834c-4bf9-b0cd-c2dc6fad3dbf": dm01.SzubsKinTwilightGuardian,
	"a7eceb07-4f6d-4b2b-8dba-7a3df8f803f7": dm01.StingerWorm,
	"edd6cffc-8c91-4682-b8af-64cfe823103b": dm01.SwampWorm,
	"a8503655-fdcb-48e2-bfb0-0ad3aae31f0e": dm01.RothusTheTraveler,
	"e2e5e1ef-c613-449a-8400-15581082501b": dm01.CoilingVines,
	"bee69327-ca6b-455c-b3dc-463fc3284b61": dm01.PoisonousDahlia,
	"bbc655b3-3676-4cda-9554-e2d465e20b99": dm01.ThornyMandra,
	"c971ff15-5735-4d61-bb16-c805130ca405": dm01.HunterFish,
	"446eaf96-36c8-4093-b4b2-e77e7afb6e3f": dm01.Seamine,
	"c9b98336-312d-4d08-8add-6b820a88815f": dm01.UnicornFish,
	"dbdbad44-6a62-4eff-b8f1-95f56588a13a": dm01.VampireSilphy,
	"f3ded71d-3cf9-415b-a9d2-b759ca0ce07b": dm01.BloodySquito,
	"dd9d1cc1-01cb-4bf9-80c4-821e3c449887": dm01.DarkClown,
	"7b22cc2c-3a4a-4f50-9e61-fb9646a762cd": dm01.AuraBlast,
	"7f225860-af37-47ac-9b36-1480872576b6": dm01.BrainSerum,
	"5cafc789-e730-4472-9a62-8b333b2691e6": dm01.BurningPower,
	"71d90484-c144-4dcf-ad8e-23e7e55f0f2e": dm01.ChaosStrike,
	"452aead9-2a65-46b1-84b2-383fe99ddc5f": dm01.CreepingPlague,
	"87a102b5-71fd-410a-a8f0-c35182217f08": dm01.CrimsonHammer,
	"2c8ded77-89f3-4625-aa3e-6576b83e0384": dm01.CrystalMemory,
	"6f2cc530-1228-4b03-9ec0-ba24f6a367bf": dm01.DarkReversal,
	"d1703c3b-8e49-4959-8322-ae11a7ca6632": dm01.DeathSmoke,
	"32acfe8b-90fc-4ba9-b6ad-7655c0abee12": dm01.DimensionGate,
	"ae797f95-54b1-48e9-9216-f315b39826bd": dm01.GhostTouch,
	"0ec572b0-ffaf-4abd-a540-ba26c98aacc5": dm01.HolyAwe,
	"95bfccf9-91cf-4ab9-8298-c95bc368bf0b": dm01.LaserWing,
	"35a9315c-2c08-46e0-b96b-daf3e8e996ce": dm01.MagmaGazer,
	"b12f1d66-46ee-49b9-878d-59cc3d515633": dm01.MoonlightFlash,
	"a2e11f7f-63f4-4357-9e77-5314576dff45": dm01.NaturalSnare,
	"858b2c1d-5507-46f0-8840-151ee59f760b": dm01.PangaeasSong,
	"9430e127-ce64-4572-b386-f1ce3f50e94a": dm01.SolarRay,
	"c3389188-718d-45bb-8946-0a572d96916b": dm01.SonicWing,
	"be80f0e8-05b7-4914-916a-f24d5e616ea8": dm01.SpiralGate,
	"a3331db6-d3c2-4b3f-9fcc-f4aa9fc2bb52": dm01.Teleportation,
	"5883180e-d88c-4f24-b17c-f5a837420147": dm01.TerrorPit,
	"48c5c29b-2f4e-4a57-86b4-864c6f0dc124": dm01.TornadoFlame,
	"be8c0d0b-dcab-402c-8e7b-878e35bacca7": dm01.UltimateForce,
	"68d78fd4-db8a-43a6-8eb6-e1435cfc2959": dm01.VirtualTripwire,
	"c62208ec-efc8-4b08-bb01-6cd5251b0969": dm01.BlackFeatherShadowOfRage,
	"e2b992ee-91a3-49d3-8228-7be60a0b9ec5": dm01.WrithingBoneGhoul,
}

// DM02 is a map with all the card id's in the game and corresponding CardConstructor for dm02
var DM02 = map[string]match.CardConstructor{
	"40439f79-8f48-4e62-9009-cb06798ef7ac": dm02.EthelStarSeaElemental,
	"48ab3f2b-4ae3-41a4-ae6f-61b49c958bdb": dm02.BarkwhipTheSmasher,
	"0bea1262-311a-47b1-888d-dd065cfe3d7f": dm02.EngineerKipo,
	"0dca6f6c-c426-4c88-b283-043527f04bb3": dm02.FighterDualFang,
	"1eca6a24-9270-477f-a588-80859481ef94": dm02.SpiralGrass,
	"2c7e38e1-0546-47ab-9388-383d093405b2": dm02.FortressShell,
	"3f0fb8f6-d01e-4005-8340-b84584f50a2a": dm02.CrystalLancer,
	"4b715b5c-2e82-4686-9c9f-4ce1e5503621": dm02.BurstShot,
	"05d946f7-5977-4f51-8bca-ecb39845f1a2": dm02.BolzardDragon,
	"5cf64846-0eb2-4e8d-bf15-4ca573f96e58": dm02.KingNautilus,
	"5d095b28-262e-454d-96c7-9174ed83e3f6": dm02.LadiaBaleTheInspirational,
	"6e381955-231b-4e4e-a14b-82509a5e193b": dm02.RumblingTerahorn,
	"9fed2257-362f-43c7-b50e-5526ccf799aa": dm02.MiniTitanGett,
	"17ee5046-c3fd-4422-af14-c54a4be8d9a2": dm02.LogicCube,
	"39b8b1c0-bc1d-445b-b60d-91cabfe62fb5": dm02.DarkTitanMaginn,
	"41e8d5c2-bfeb-48bb-ab9e-aaee79852c89": dm02.CrystalPaladin,
	"60d8c6a6-20c1-425c-9ecc-b56981a70e21": dm02.ElfX,
	"66ac493f-b836-46c9-a28e-09e7bd040064": dm02.Gigastand,
	"66ee3b9e-fdbf-41c2-9363-5327572706f2": dm02.ChaosWorm,
	"84b7a5f2-cbfc-4d2d-a757-6685aa38c241": dm02.GrayBalloonShadowOfGreed,
	"94e5d2a3-d8b2-4fce-9899-61adb063dc60": dm02.SilverAxe,
	"96ecef8e-0485-418d-9d0d-a4169c3b70b7": dm02.FonchTheOracle,
	"98b1afbe-5a0a-461e-800c-34d02339f21f": dm02.ThoughtProbe,
	"215c4cfb-2a22-4ee1-b4ea-28ac24a1eeee": dm02.UltracideWorm,
	"733a4f35-7470-40e3-9cd7-479aa965bfbb": dm02.HorridWorm,
	"749d1209-4847-4401-8d70-8c95fd26c220": dm02.MagrisVizierOfMagnetism,
	"839f780b-067b-4b3c-9090-7cd614375da9": dm02.AmberPiercer,
	"1960f3c2-5321-4da1-be56-50a7e98cae0d": dm02.PhalEegaDawnGuardian,
	"4916f651-6169-4e8b-be5d-6a3237dc0a70": dm02.MarrowOozeTheTwister,
	"7956b4f5-b910-403d-b388-b67c837b7e99": dm02.ScissorEye,
	"9311b3ec-fe9e-45bd-a989-dc00aec3ca60": dm02.DogarnTheMarauder,
	"24353d06-89ef-4867-9513-485750d01e10": dm02.ArmoredCannonBalbaro,
	"99791f19-21ee-4ee3-a6b0-e26702c2a380": dm02.CriticalBlade,
	"5781986a-4ca8-48e1-a5a1-95e820455bce": dm02.HypersquidWalter,
	"6406411f-e9ae-4b55-8213-7e18c4cd9aee": dm02.StainedGlass,
	"9275747d-f2bb-4298-9b70-7075b17d1e0d": dm02.XenoMantis,
	"206c53d6-8068-4fec-bc5b-287de58eed5f": dm02.WynTheOracle,
	"a219c895-5822-458e-862f-6ade1529c42e": dm02.LeapingTornadoHorn,
	"aa1cca8a-3140-4180-9c9f-3f126dd16a68": dm02.ReconOperation,
	"ac636f90-215f-4934-bcd3-abfec69a5fda": dm02.AquaBouncer,
	"b85b0492-132c-46ac-85e6-c8b5b466598a": dm02.LagunaLightningEnforcer,
	"b691251e-8b9b-4436-8ca8-d9d3aba3ac72": dm02.ResoPacosClearSkyGuardian,
	"ba8c4d25-a352-4725-a3aa-6eb015f99d6c": dm02.Galsaur,
	"bf5ce598-f604-4c55-b0c9-9ed0e059aeec": dm02.ManaCrisis,
	"c125f786-e6d5-4477-8ab0-1e92d6eed348": dm02.CavalryGeneralCuratops,
	"cc5c643e-5120-421d-b26f-76be381dead7": dm02.RumbleGate,
	"cf536a4d-1512-44d8-b0e6-615d0725e3f0": dm02.MetalwingSkyterror,
	"d7e4a7da-6131-42e9-adf1-de0178549bd5": dm02.SilverFist,
	"e1e112d7-11e1-4f01-9c91-00a2b1626043": dm02.Bombersaur,
	"eac1bc57-bdf8-4629-86b3-9609d1bf2aba": dm02.ArmoredBlasterValdios,
	"ec3213d2-6914-4091-bf86-869354a36b15": dm02.PoisonWorm,
	"ecd61b34-85ca-4a91-b993-7a8c9976a2b5": dm02.Corile,
	"f4a364f5-d0e9-4777-b51e-6dc6e39b803c": dm02.AquaShooter,
	"f9b47778-c7d5-48da-aac3-1ddf8efbbbf1": dm02.PlasmaChaser,
	"fae765a5-fa73-4627-94d1-f74f2ed00792": dm02.LostSoul,
	"fcf2f484-c471-4a5d-bc41-1fcd56604d73": dm02.GeneralDarkFiend,
	"fd9f6d26-8fd8-43f2-84e5-b83493a1106e": dm02.LarbaGeerTheImmaculate,
	"ff33e31f-4f1c-4c07-b293-2e65a13e1077": dm02.RainbowStone,
	"ed3931f2-2b60-48d4-b82f-361d3a8c5bd7": dm02.EssenceElf,
	"a73bf1d3-b4a6-48cd-89d3-1983f160304d": dm02.DiamondCutter,
}

// DM03 is a map with all the card id's in the game and corresponding CardConstructor for dm03
var DM03 = map[string]match.CardConstructor{
	"70270aa3-ff24-476c-be22-5b9f48fc682a": dm03.MiarCometElemental,
	"9c7e3304-3aff-4362-a687-b5ca5333fe98": dm03.ChaosFish,
	"9df4d8ac-3c86-4c1f-b916-c9a384b0340f": dm03.GirielGhastlyWarrior,
	"41c2e4dc-460f-459a-b7cf-ef17b5c9a4eb": dm03.GarkagoDragon,
	"54652ec5-3bc3-4124-9c67-b05ba56def5f": dm03.EarthstompGiant,
	"4459f97b-8927-4231-a11d-f4f88175b71c": dm03.AlekSolidityEnforcer,
	"9d9468f8-3eb3-448c-8638-5d77c77f936b": dm03.AlessTheOracle,
	"b275fbf0-5355-45ec-b3a8-a956cf898ae6": dm03.BoomerangComet,
	"6b6ff714-0602-48da-9ee0-2578aa8ea32e": dm03.LenaVizierOfBrilliance,
	"96370266-e077-47c1-968f-0184bcb94227": dm03.LogicSphere,
	"b052351d-dd6c-4c40-a6d7-840e23398ff1": dm03.RaVuSeekerOfLightning,
	"eb46cc3f-75b9-4b50-851f-bf1fd3e76c06": dm03.RazaVegaThunderGuardian,
	"c49be3ab-da9b-4af8-8377-0595ca3160ce": dm03.SiegBaliculaTheIntense,
	"b9b1a83a-626c-471a-82ff-3e2372c9d838": dm03.SparkleFlower,
	"a097478c-af87-4316-b42d-888418d07919": dm03.SundropArmor,
	"7fdf0d01-a999-4f31-8a37-945396796998": dm03.UrPaleSeekerOfSunlight,
	"f76fd433-24ba-41e3-81eb-622cad7247e2": dm03.AnglerCluster,
	"23a97f92-d53d-4b9f-8e33-ca090bd30c1e": dm03.AquaDeformer,
	"d96cdbba-0e60-4d8e-9394-4830b11f559d": dm03.Emeral,
	"4852b823-7b96-4669-809b-edc14c042567": dm03.FloodValve,
	"39e0b4b6-eaa5-4673-9538-aa9d2c7fc5d8": dm03.KingNeptas,
	"2f2dae63-bb68-4d89-b7b1-48dd14eab40e": dm03.KingPonitas,
	"232a9fdd-e289-454f-9593-5b766c969a1e": dm03.LegendaryBynor,
	"e14514a5-592b-40ed-b693-7cb8f649d93a": dm03.LiquidScope,
	"c4f3db71-e9e7-4ab6-95fc-1174b4f42aca": dm03.PsychicShaper,
	"0f9a0c09-1d5b-4162-94ca-5399f8f37d63": dm03.Shtra,
	"fb3a5abd-c16a-4365-98c0-79753fdcd91d": dm03.StingerBall,
	"e4ad6b63-5c97-4002-9b0a-b38ba8b5312b": dm03.BaragaBladeOfGloom,
	"dae59bf4-cb23-40d0-aed2-bdbf953f73f5": dm03.BonePiercer,
	"826de4ef-a10d-410a-aed1-066783babbdc": dm03.EldritchPoison,
	"b3975c0b-2978-4b1a-8225-78d420ff941d": dm03.GamilKnightOfHatred,
	"21a9e35f-4219-490c-8963-3e1fee66c71b": dm03.GhastlyDrain,
	"03b16fba-71e7-4328-a605-4643bb5cfbf8": dm03.HangWormFetidLarva,
	"12d21399-f499-432b-bbb0-8ca1088b33c7": dm03.JackViperShadowofDoom,
	"ec9ebf80-3060-48cf-888a-bd00b61b3746": dm03.Mudman,
	"c08d8e5a-9f4b-4fe4-bf20-d5c7fca9ea15": dm03.Scratchclaw,
	"f773e1f2-37d0-45c2-a34f-63e91094aab1": dm03.SnakeAttack,
	"bfef8796-4214-4cfe-af86-4a8608a6fd7f": dm03.WailingShadowBelbetphlo,
	"b6109c08-3d01-486d-b544-5dfc11d341a0": dm03.ArmoredWarriorQuelos,
	"479b477f-f535-4564-ac0c-5f0aeaafe914": dm03.BabyZoppe,
	"d74a5d6f-a24e-4ecb-a163-7c696be8c06c": dm03.BlazeCannon,
	"e27ac147-3d7d-42d7-a3a4-2e3a1eccdb2c": dm03.BoltailDragon,
	"0f5a2f7e-74df-4ae5-ab0e-300b1761b417": dm03.ExplosiveDudeJoe,
	"caca44fe-8e47-4863-b1e9-7c7c09eb7ddc": dm03.Flametropus,
	"89bf69b3-25a7-4f9f-884a-fb9e7dd93efd": dm03.MuramasaDukeOfBlades,
	"d2cc5ca6-c0a3-4a83-935f-08d0e0878983": dm03.SearingWave,
	"cf6673ec-4bae-4891-9c80-e3e3a8269914": dm03.SnipStrikerBullraizer,
	"5646364b-4018-4556-9b20-265ef3fb372d": dm03.UberdragonJabaha,
	"0a5600c5-65a0-439d-b19c-4bf5ee036e32": dm03.VolcanicArrows,
	"9cff685f-3491-400b-b9d4-43c383193dc7": dm03.AuroraOfReversal,
	"22dae49f-f802-46d1-ad1e-400b980b80ab": dm03.DawnGiant,
	"e202fcb0-08a1-46ee-8654-66406ea436ce": dm03.Gigamantis,
	"5170634d-1574-4a06-b3bd-3c546fb9191e": dm03.ManaNexus,
	"f200f25c-479b-4381-967e-590bedd9dc90": dm03.MaskedPomegranate,
	"c3c09eac-d9b9-4eb6-8056-94afa4ba3927": dm03.PouchShell,
	"64371ab9-a8bc-4ec9-9b9b-f190c8ac1001": dm03.Psyshroom,
	"e8910059-0471-44ac-a573-a3fe8ed207d9": dm03.RagingDashHorn,
	"fc25a99c-7ef9-44bc-9204-ca84f6248116": dm03.RoarOfTheEarth,
	"8dab06d5-93c3-4ac7-abfb-3ce2e63d77e5": dm03.SniperMosquito,
	"84595f9f-c9c1-4d50-8e8f-29e5ef63bfbf": dm03.SwordButterfly,
}

var DM04 = map[string]match.CardConstructor{
	"0ce1bba5-376f-4c9b-9865-730f9e62a7d7": dm04.FullDefensor,
	"2e5452fe-1225-401a-b98b-b8a08639a467": dm04.SkeletonThiefTheRevealer,
	"2fe88eac-54da-4e49-a6b3-22b1b8c7da65": dm04.CloneFactory,
	"81ebd0e6-94dc-44e6-958a-6df15dae091c": dm04.GalklifeDragon,
	"70ed7fd3-afb9-4b62-9c3c-6b26b8eccf04": dm04.DoboulgyserGiantRockBeast,
	"746d520f-7486-4eba-ae59-58473092a999": dm04.BlastoExplosiveSoldier,
	"bc430287-5cac-42b5-a6f8-5edcfe5cea58": dm04.ChaoticSkyterror,
	"0103d28a-1c07-4cc3-916e-1fd67ef9595a": dm04.KamikazeChainsawWarrior,
	"6567373c-9549-4dc4-a7b6-824c6d7c4e84": dm04.Magmarex,
	"e6c76df1-24c8-4125-9f9f-8ae3b2bc61f6": dm04.MegaDetonator,
	"43cde867-4cad-4ed7-867e-b01d1577b571": dm04.MissileBoy,
	"1484ec6d-c1b5-4fc4-abaf-a16c08cfc5f7": dm04.PippieKuppie,
	"4dd64294-fe68-4337-b6fc-48e1a473a1ce": dm04.SwordOfMalevolentDeath,
	"7b6fe59d-8b11-4f8b-9a0b-1ce22673c274": dm04.KingAquakamui,
	"187a6327-b0e0-42a3-8cae-9293b41927ac": dm04.AstralWarper,
	"a31acf59-bf6a-4e52-9158-53103e5e9951": dm04.AquaGuard,
	"a9469dd3-73fa-4a8c-a9dd-7b852487308a": dm04.AquaJolter,
	"4143565f-a7a0-48f9-9601-67ad15ca0933": dm04.HunterCluster,
	"831f5090-667f-41e3-863d-2db3d8be68db": dm04.HydroHurricane,
	"979c8fda-550e-40af-a88e-d26bee3cbbb4": dm04.KeeperOfTheSunlitAbyss,
	"8155edc9-5ab2-4295-95e4-2cd07c438e61": dm04.Marinomancer,
	"e4408b19-a426-4cc3-8b81-281499e80a88": dm04.SmileAngler,
	"b80de483-5aa7-4181-8672-bb3a5c11b438": dm04.NiofaHornedProtector,
	"38e927bc-74a1-498c-bbaa-3443999ac7a0": dm04.SupportingTulip,
	"a04666a7-a12f-4777-b4f9-420f7ad4b253": dm04.AncientGiant,
	"09604e37-5e09-4042-b5cb-5c180f048167": dm04.CannonShell,
	"206c032c-2b4a-4278-bfac-af9ff8ece118": dm04.DewMushroom,
	"7e9369a0-f137-4e37-8576-4733993c2ba4": dm04.ExplodingCactus,
	"ae56a6c7-b118-44cb-bb6d-c45a9484c598": dm04.MysticInscription,
	"e55c0de2-8366-4585-ac08-ebaea0e96818": dm04.SwordOfBenevolentLife,
	"e2305d22-7105-44a0-b1c0-5c9caba974be": dm04.ThreeEyedDragonfly,
	"aa5deaa8-4935-47ee-a6dc-a036398886a1": dm04.Torcon,
	"f2103ae0-1ea6-411a-b6f8-553f995cb65e": dm04.BallomMasterOfDeath,
	"8e10db1c-1445-45c0-b7d2-3b888ed87998": dm04.TroxGeneralOfDestruction,
	"9182f914-2bd2-41e6-b950-9edb2035aff1": dm04.ChainsOfSacrifice,
	"a646e98b-70ce-402d-b995-3678757c86d3": dm04.Darkpact,
	"79e32904-a0e9-4596-99ad-76bda92e672c": dm04.Gigabolver,
	"70f08780-cf4b-4c4a-8188-5235157ce863": dm04.GregoriaPrincessOfWar,
	"bd3d2c66-b38a-40eb-9333-3509e0e213c9": dm04.GregorianWorm,
	"ba955ab0-5bb3-4aaf-82f3-293522e65a9c": dm04.Locomotiver,
	"593ca472-b8d1-4fec-ae07-0102afd12a48": dm04.MongrelMan,
	"b0b950c5-5498-4edc-a7dc-14991961fd2e": dm04.PhotocideLordOfTheWastes,
	"3bd6cfa5-7641-441d-a76a-b6df468ec292": dm04.PurplePiercer,
	"da8d62d3-47b9-450e-961c-9bcc171ff0ee": dm04.ShadowMoonCursedShade,
	"a41ca346-193d-4c2f-bccb-382ca2591120": dm04.SoulGulp,
	"da7988b3-43dd-48fb-9339-b1834b61d1ee": dm04.VolcanoSmogDeceptiveShade,
	"d01f8146-c033-4150-a18e-b648c6786d77": dm04.RimuelCloudbreakElemental,
	"7d4b64b0-1672-47d4-b54d-1758b0bb08cf": dm04.AlcadeiasLordOfSpirits,
	"6a2f2777-e7be-49ee-8bb4-d963f1ecd409": dm04.AerisFlightElemental,
	"5da6bf2e-9de1-4166-ade0-a9ded5b99612": dm04.AmberGrass,
	"d74da0b7-3c0d-4e27-85c9-4ccf0df621d1": dm04.FuReilSeekerOfStorms,
	"55c0e0e4-f585-423b-9b23-054156ab3de3": dm04.GulanRiasSpeedGuardian,
	"0993f972-634e-4175-86a5-25074a5bd165": dm04.KolonTheOracle,
	"7a94b69e-335c-489a-8c5d-1294a9cb960f": dm04.MilieusTheDaystretcher,
	"6a909066-2011-4de7-9a97-cab333dca4cf": dm04.MistRiasSonicGuardian,
	"339ba23a-0228-4677-a90a-c7efc843ee73": dm04.OuksVizierOfRestoration,
	"ae43a5d8-8bae-4e0b-b31c-66a37c94c719": dm04.ReBilSeekerOfArchery,
	"ebb2a7f2-ea2d-427f-8dca-0a650553736d": dm04.SariusVizierOfSuppression,
	"afa724a7-6a5e-425d-984c-b1a977b0c7f8": dm04.ScreamingSunburst,
	"d48b24ad-2b6c-4a00-affc-f95be66976fc": dm04.WhiskingWhirlwind,
}

var DM05 = map[string]match.CardConstructor{
	"1fe34153-7cd0-40e9-bbcf-96dd7074860c": dm05.TwinCannonSkyterror,
	"82e062b1-bf9d-444e-9cd3-2fdef007ba64": dm05.BolgashDragon,
	"15456751-98e0-4cbf-b2e9-649427bcf8ae": dm05.LaGuileSeekerOfSkyfire,
	"a8189c59-1f1b-498d-a289-113ce02ab631": dm05.SeaSlug,
	"ab743952-7857-490a-b6d6-6ecd9f0d68c8": dm05.RikabuTheDismantler,
	"d98448c3-fd04-452b-a338-c9674c69e96f": dm05.CannoneerBargon,
	"dac08511-6c9a-44cf-b328-237acb678305": dm05.BombatGeneralOfSpeed,
	"f78a43be-b874-4a05-b681-48647fde9e48": dm05.BillionDegreeDragon,
	"3e6d2002-39f0-4be9-bad0-1738a299731f": dm05.BallusDogfightEnforcerQ,
	"3e63d487-a25c-405d-8c04-df8c5e414fd6": dm05.SplitHeadHydroturtleQ,
	"4ac7c097-cdc3-42e8-9203-1a481274f9d2": dm05.BladerushSkyterrorQ,
	"5ad88bc8-5991-46d9-a98a-19c7569a4f05": dm05.RuthlessSkyterror,
	"5f5a3a3f-3675-4e5d-8bd6-0c76cb8c4ce4": dm05.DeathCruzerTheAnnihilator,
	"cf0c049f-f047-4c57-9043-7911eb28395d": dm05.EnchantedSoil,
	"f0587622-589f-46c5-9857-d73ee9db4db4": dm05.SkullsweeperQ,
	"4353b917-261d-4365-9f49-dc2d52f6a8bb": dm05.AvalancheGiant,
	"5ff3c63c-30df-4f5c-acce-796f5b6c2dac": dm05.SmashHornQ,
	"6b3ec428-4fa1-4646-ad49-54fa379df2e7": dm05.NocturnalGiant,
	"7a59d73d-f340-482c-beb3-55ee83aa6222": dm05.MoonHorn,
	"8e3510d7-fe27-4e62-91b7-b6c869a615e3": dm05.SchemingHands,
	"8fac6848-e179-4611-a56a-f76f7716f0a8": dm05.LurkingEel,
	"9b9446d1-cda6-4d57-91d9-9cf4264e3fd8": dm05.CrowWinger,
	"22fc691b-58bd-49b4-b807-8b0592b8c8db": dm05.SnorkLaShrineGuardian,
	"27eca314-2540-45e0-8201-71e33bb14da7": dm05.CyclonePanic,
	"50d6a40e-dfaf-46e1-81e9-376f9aa8ffbc": dm05.GlorySnow,
	"81e05a20-875a-4047-af47-9850dac9a863": dm05.SyforceAuroraElemental,
	"372ecf37-071a-4ddd-bf50-b27e7fcb5913": dm05.SyriusFirmamentElemental,
	"680f0ae7-9962-4dbf-88da-b1faf3e1e066": dm05.CalgoVizierOfRainclouds,
	"c7fec5e8-4e56-451b-a7b6-ad08680703a4": dm05.LaByleSeekerOfTheWinds,
	"d3a108f2-57f3-4fba-9b1b-b85a5ca788ed": dm05.KulusSoulshineEnforcer,
	"117765de-dbde-45ae-941d-2b864421ba4c": dm05.KingTsunami,
	"a9e9a9c5-9125-4116-8fe9-f4087615b8a1": dm05.KingMazelan,
	"f5b9c2d5-970d-4578-a521-acccd856cfdf": dm05.AquaSurfer,
	"870e7b00-eb00-4546-9f15-a5cb6cb976d9": dm05.VashunaSwordDancer,
	"42949cc1-d2e4-459e-9ccf-ac93ebe636de": dm05.SlimeVeil,
	"f288e855-44bb-4649-8c9d-de4a61fd228a": dm05.BrutalCharge,
	"aa5799bd-aaee-4ff4-89ae-374db57cfa83": dm05.MiracleQuest,
	"69164e5d-2c0e-4369-b0c2-0494840128fd": dm05.DivineRiptide,
	"826eecb2-0beb-47f0-afb8-2d01bd51491d": dm05.CataclysmicEruption,
	"39fee52e-e539-4bd9-99b7-03e61b4afea3": dm05.ThunderNet,
	"f6ff8845-afc0-4958-8673-fad12058193a": dm05.BloodwingMantis,
	"4246218b-c66b-4f8f-b05a-1e7bff9f39a5": dm05.ScissorScarab,
	"5145a445-45ce-430c-ba58-c15c5e8af984": dm05.AmbushScorpion,
	"75b315ab-c169-46b0-a693-1eb52eef24f8": dm05.ObsidianScarab,
	"89bdc9a2-72c8-4ce7-afd4-626538ed4c82": dm05.SolidskinFish,
	"ece74930-3f8b-4d35-95ed-94639dd251cd": dm05.SpikestrikeIchthysQ,
	"a422a702-fc94-45e8-b6d9-aa6b92e8797f": dm05.BlazosaurQ,
	"c0c9c3c6-6b8c-49a3-b6ff-ad9d7c2bf5a7": dm05.GalliaZohlIronGuardianQ,
	"d5a09c99-985e-412e-b77f-558c059db7ee": dm05.JewelSpider,
	"fd297f8a-c5ae-46c3-9d59-a13de5736b0e": dm05.HornedMutant,
	"1053c39a-68c3-4811-9c7c-049914f420e3": dm05.KipChippotto,
	"a8f4d303-72c0-4b67-b6af-d99d267de0c5": dm05.BalloonshroomQ,
	"4c5323af-2c14-4bc7-9cb8-7b967c7b8586": dm05.Pokolul,
	"f66ca3b3-7300-41ef-8319-9de6dfc4e39a": dm05.SinisterGeneralDamudo,
	"38db41d2-2a6d-4f6e-83e8-265208be45ea": dm05.Gigazoul,
	"60b50334-f4c9-46dd-aa87-76f4e67cb5f8": dm05.LeQuistTheOracle,
	"181126d5-0ad8-4033-a9d2-c8ca939f3fac": dm05.SteelTurretCluster,
	"1064bbcc-0f15-4757-8d88-93084edaba3f": dm05.Gigakail,
	"cfedb8c9-419a-4c0b-b05b-a3a30436f4b0": dm05.WispHowlerShadowOfTears,
	"88cf0a1e-5506-4d5a-a2fd-0d1bb60f6abd": dm05.GigalingQ,
}

var DM06 = map[string]match.CardConstructor{
	"00f39646-e15a-47a1-a085-c3522bfa5146": dm06.RaptorFish,
	"187d4a6c-9b92-4c7c-8507-e1de185bce02": dm06.CrystalJouster,
	"3d32a70c-303d-46f8-862a-f7d5a3cfaafc": dm06.NeonCluster,
	"4269337e-5772-4d22-9474-e3068cf21de0": dm06.UltraMantisScourgeOfFate,
	"a5fa55a7-8ded-4e8b-9b45-f791427cffd8": dm06.CrazeValkyrieTheDrastic,
	"0277a21a-7456-48f0-84b1-fbe7f5010196": dm06.CutthroatSkyterror,
	"d92b07a3-5dba-448b-94c6-310a588956cc": dm06.CursedPincher,
	"9d7de648-c25f-4439-835b-359147cca717": dm06.JunkatzRabidDoll,
	"42ee6087-94cf-45cc-9e0b-baf73a008d81": dm06.LupaPoisonTippedDoll,
	"708d2e3a-76dd-40d6-9278-345383b3047a": dm06.PyrofighterMagnus,
	"f2344e6e-0ee3-4f17-8984-b01025db7e11": dm06.BazagazealDragon,
	"b42f1258-714b-4051-bd4c-2013e40a4c6b": dm06.ValiantWarriorExorious,
	"03e90b61-de6b-4512-a656-2832968b694a": dm06.AutomatedWeaponmasterMachai,
	"1768f54e-5855-45d9-9c6e-306fc2926451": dm06.MightyBanditAceOfThieves,
	"78d73122-a35c-448e-871f-8bf713d81fb7": dm06.Gigagriff,
	"a441c54e-8edc-4c70-b9e0-64fe73a468ff": dm06.CarrierShell,
	"8eca8a5e-f039-4103-b113-5d8c5566a0c0": dm06.SlumberShell,
	"52415948-e965-4636-a240-793f8c29de7c": dm06.CharmiliaTheEnticer,
	"41b82092-4d34-4b3c-934a-d7bd4fca654a": dm06.GarabonTheGlider,
	"1b83ff2c-3095-40d5-bd76-e1566836a8ea": dm06.RikabusScrewdriver,
	"27c58f10-82ef-47ca-8a69-1cfa2057743d": dm06.PicorasWrench,
	"c0dd4b57-b334-4954-88ff-8a791df67b28": dm06.VileMulderWingOfTheVoid,
	"de383025-127a-445b-b705-3f258ac6cccf": dm06.ZorvazTheBonecrusher,
	"6a525afa-ec9c-44c9-93dd-31e76d82538d": dm06.DaidalosGeneralOfFury,
	"b88a3f25-12fd-4c25-8bee-6f1e81fce51b": dm06.GnarvashMerchantOfBlood,
	"5006e5ef-2b00-434e-94ae-9d992d0df5c1": dm06.LoneTearShadowOfSolitude,
	"ed3d85f3-f1a4-41b7-88ab-6ecef15948b7": dm06.GrimSoulShadowOfReversal,
	"edb47886-0edd-4aeb-867b-f4d71bb626dd": dm06.FrostSpecterShadowOfAge,
	"12601fae-c2b7-4211-9f84-0922628d384e": dm06.SkullcutterSwarmLeader,
	"8805e361-8971-4251-9bd5-a1f044be4185": dm06.GrinningAxeTheMonstrosity,
	"b2ad3f06-2901-498b-ab40-4e37f2b3649d": dm06.SplinterclawWasp,
	"93ef36cc-4426-46e8-8e12-84e8435e6892": dm06.TrenchScarab,
	"59da70cc-e655-4e70-b467-8024dc7c472e": dm06.QTronicHypermind,
	"305bb0c5-7bfe-47cd-a83f-80d8cb620c8c": dm06.HazardCrawler,
	"7372e3b8-85d1-4d3d-a704-be821ac2211a": dm06.MidnightCrawler,
	"c1a29dac-6649-4920-bc4d-b0b8070ed15f": dm06.ThrashCrawler,
	"6bfebd3d-64ff-4321-bb0b-e994ca8f811e": dm06.TelitolTheExplorer,
	"da51845c-4a6b-4c36-9c7d-fbb654ba2aa2": dm06.KanesillTheExplorer,
	"0743092b-fd5e-41ab-8e38-d6664c1741cd": dm06.MadrillonFish,
	"47b35edf-3982-48af-b98b-5548300de386": dm06.SchukaDukeOfAmnesia,
	"74ce4102-f6af-467e-8d81-464ac9ebe25e": dm06.Sopian,
	"2982c4fd-1518-4686-b523-2317600ab425": dm06.InvincibleAbyss,
	"3d6b7853-c3eb-446c-aee1-984ed00b04ac": dm06.PhantomDragonsFlame,
	"53ecc374-d627-43b0-8091-220c313de700": dm06.InvincibleTechnology,
	"5fc2d880-b508-460f-b1d3-e2d9a7377612": dm06.SphereOfWonder,
	"646245a5-3c70-4842-912c-82a288846f97": dm06.SpasticMissile,
	"6dfdd98c-da99-4a5d-9595-51430026d97d": dm06.FaerieLife,
	"79012009-cb97-4e3a-96b5-76fbeb37ea2b": dm06.PangaeasWill,
	"a50a7fba-0fd7-49e1-887f-8d229f3a5b2d": dm06.MysticDreamscape,
	"b2a80ade-8d00-4aa7-ad89-30473dfecd1b": dm06.ProclamationOfDeath,
	"b455c70a-bcac-4063-9bc3-9793b7365b66": dm06.FutureSlash,
	"c99a26f4-ee0a-48d2-aae8-89a775eb208a": dm06.MysticTreasureChest,
	"d8aee1e9-0799-4acf-af2a-3a8fd1b4eb8e": dm06.InvincibleCataclysm,
	"e56c04ae-1edf-4308-87b6-697e0d9a85d4": dm06.InvincibleUnity,
	"f1bd7b39-7582-4f80-ba4c-423ef16f7a96": dm06.InvincibleAura,
	"cf8b1a10-a425-4670-bb50-af7c4fb132d6": dm06.BondsOfJustice,
	"b7f236fd-e7eb-41cc-912a-5239c134f265": dm06.EnergyStream,
	"1674d51c-378d-4257-b2b7-e587d1edbdd8": dm06.ChoyaTheUnheeding,
	"65df0a97-9686-4a93-a586-28b491592806": dm06.ArmoredDecimatorValkaizer,
	"791d2be3-e30d-4215-8db7-aaccda3dc890": dm06.MigasaAdeptOfChaos,
	"5e73baca-4680-4c67-aa52-e4da9ac44e15": dm06.LegionnaireLizard,
	"9ff648e7-b909-4df7-ac27-21f57dd71f59": dm06.BadlandsLizard,
	"e415e7ee-ab90-43ec-9443-05dcbb7f6b88": dm06.BallasVizierOfElectrons,
	"ab0c0696-3abf-4a10-ac41-805a6c538849": dm06.ChekiculVizierOfEndurance,
	"42af9ccd-021c-48e5-833d-d7b6a08676d2": dm06.ChenTregVizierOfBlades,
	"b900ea21-4caa-44b0-81ce-b2dbe6a06722": dm06.SteamStar,
	"7cf434ba-e27e-4674-b927-df22fa053a26": dm06.RippleLotusQ,
	"4717faef-1065-4153-a509-854c22637e27": dm06.BolmeteusSteelDragon,
	"57b5a157-2ccc-40a7-b8d4-b1b738071562": dm06.LaveilSeekerOfCatastrophe,
	"aad5a2ae-0ad3-4ba1-b620-47913c2b3c7d": dm06.DavaToreySeekerOfClouds,
	"ca1f21a4-7f8e-45f3-b7e3-467bdcbd3b10": dm06.BlissTotemAvatarOfLuck,
	"d146c950-37d9-4893-95ce-30bd237b5097": dm06.ClobberTotem,
	"adc7eb5d-862d-430c-bd7d-ca51f0c94b02": dm06.QTronicGargantua,
	"e2e64a5c-ec8d-464a-bbe3-bacbb01da399": dm06.LightningGrass,
	"7c33ae39-5079-471b-982c-98da274c9892": dm06.RazorpineTree,
	"cc7916d4-9f3a-4993-8283-134f3bba105c": dm06.CliffcrushGiant,
	"e0538282-a50e-4097-840c-80f262ce0416": dm06.CantankerousGiant,
	"eaa15452-23bb-4fad-a557-9df3c9752292": dm06.RumblesaurQ,
	"fee203a9-b5a8-4d01-abb9-12d07c62475d": dm06.CosmogoldSpectralKnight,
	"efa301fb-a605-4ab5-ae3d-1fb49e2a83e6": dm06.MoontearSpectralKnight,
	"f0ac300b-2ebd-49b6-aa43-34d5a131a715": dm06.GarielElementalOfSunbeams,
	"3031a721-6725-42f3-a694-08e931eb9fc6": dm06.FeatherHornTheTracker,
	"8fdf84b5-7fa2-4caa-820e-5c045fc51156": dm06.ParadiseHorn,
	"eb15fb12-f761-44d7-a91f-4f85a9d27814": dm06.Zepimeteus,
	"ec0b5fbc-3a94-4f45-80f8-063daacb62e0": dm06.Aeropica,
	"d176b30a-cac6-4249-a78d-18f34b97546b": dm06.PromephiusQ,
	"63047135-feae-43bb-8610-8ce6dbec0456": dm06.GraveWormQ,
	"3370af08-dfd1-42a7-8f92-692ba7c48017": dm06.TentacleWorm,
	"ca9f7712-db3d-4233-acf4-7c98646cc1d3": dm06.RainOfArrows,
	"76da3804-df25-4773-ba2f-ea17bab89f2d": dm06.CometMissile,
	"5c424a0f-5bbd-41cd-9279-2b408f7e5935": dm06.CrisisBoulder,
	"124dc6bb-a6c3-4771-91a5-9cd2c2b198e7": dm06.FactoryShellQ,
	"0aeb82a3-2737-471d-a092-ebf02eca10ee": dm06.ProtectiveForce,
	"eca2f1a6-c8df-4abc-b0bb-41d31058ee10": dm06.IntenseEvil,
	"79f6a24a-5419-4988-bc19-b07120982732": dm06.ShockHurricane,
	"08a6f16d-36bc-48a9-a704-122875759618": dm06.Torchclencher,
	"bc8b2d5d-f595-42d2-a10c-2ee3cd336c3f": dm06.IllusoryBerry,
	"f5499baf-ded1-41c5-8386-3ac7f7c1c841": dm06.ForbiddingTotem,
	"f552c0a3-9863-4ea6-9769-3c1134da0995": dm06.ForbosSanctumGuardianQ,
	"9916970b-bb9b-4fa6-8a13-61b203e0cad1": dm06.LuGilaSilverRiftGuardian,
	"1ec9f300-c8cb-41c2-b9a8-b46408c49098": dm06.OverloadCluster,
	"cf8b1897-0044-4a13-9114-c1c3ba51bedd": dm06.AquaRider,
	"38ea8acd-e8e3-4bd9-a7c9-7d74761c3712": dm06.KingTriumphant,
	"8b7fc29b-d79c-4b08-a88e-9d055d02c6e8": dm06.CoccoLupia,
	"38617d18-b12a-4618-8a26-3effab948fcb": dm06.VessTheOracle,
	"a8dbcc5e-a9e8-4cc6-8b87-3b53a5701371": dm06.YulukTheOracle,
	"9970cdd5-bde6-4996-916e-00209a7419f6": dm06.BazookaMutant,
	"b8addd02-0a51-4445-a2c8-d7dc045c7712": dm06.TankMutant,
	"751ad736-7517-4f58-9d03-5b6afa865c84": dm06.InnocentHunterBladeOfAll,
	"f1a875ab-7edb-46d2-ac0d-31ebdcb97aca": dm06.ArmoredScoutGestuchar,
	"7732053b-6ec5-4a62-a5c1-cccff5583366": dm06.LivingCitadelVosh,
	"c30f60d5-d97c-457e-b438-e4a606136171": dm06.AdomisTheOracle,
	"c77c93ab-c973-4874-aef4-b1ef2f9614c5": dm06.ArcBinetheAstounding,
	"4121d282-d257-4b3b-8388-83fbb4829dd9": dm06.LavaWalkerExecuto,
	"7fd21958-859f-4085-acab-c736de7667ef": dm06.FortMegacluster,
	"569c34fc-614a-4aaf-a89b-d4e5dd49426c": dm06.PhantasmalHorrorGigazald,
}

var DM07 = map[string]match.CardConstructor{
	"d45c3c6a-da43-42c8-bf9f-e72aa8c59a36": nil,
	"a7179172-1a44-46af-abf9-59daf43e67d9": dm07.CosmicNebula,
	"504c57d4-d27a-42ce-bfff-36708b814c21": nil,
	"19bc25dc-2b35-4676-ab72-a8b13b1828e0": nil,
	"d9be6090-077b-4f34-95cc-d33b2266e45c": nil,
	"d1fe4c40-637e-4e5c-a769-0d2fc8aa11ca": nil,
	"c3c7ac05-c847-4842-9504-22e29fa24a2f": nil,
	"b6863540-2985-4349-aa82-05f9dbe36993": nil,
	"44293c09-875b-4582-99cd-450b544c875d": nil,
	"38504a0f-da4a-4855-9d93-08866f8b955d": nil,
	"7e664555-207b-47d6-8078-f40c5885021d": nil,
	"f603d464-3492-40ba-990d-85dc02814966": nil,
	"8501c2f8-7dd2-4c2f-aa80-8025c8998371": nil,
	"3fc5db56-616c-49c4-8ddc-2d922e9c4a54": nil,
	"a3569344-7a9a-476a-9afc-35e9c28ad209": nil,
	"119afa07-1a17-4197-8df9-3ee4254f0a58": nil,
	"4507bcbd-17f6-4035-9da5-1a6d0b41601e": nil,
	"ae65d4d1-78aa-4711-ba49-eafd98b358e9": nil,
	"7148214f-dc9c-4fb4-835c-cd6064e247df": nil,
	"e5bb03ac-0f09-44de-8fbe-b0c7f2748e84": nil,
	"0d273a22-1cf9-4f99-8321-e2282e855b57": nil,
	"7e4cdeee-0b47-4260-8ba8-3f1cf0b9ab36": nil,
	"0735c257-6554-4aeb-9aca-f3bed4022752": nil,
	"5ce98713-900d-4291-8dc4-889de1d68461": nil,
	"44c6f9ec-1832-4d62-a73e-3dab72818591": nil,
	"10048aab-7f26-4c3f-862b-199f41f132aa": nil,
	"a1ed67a1-2143-461e-b2ee-7d74ec6773e5": nil,
	"cf6ee28e-ab9b-4313-9ea9-048767476308": nil,
	"0106b420-f033-4f9e-b5a9-237f1d6ded0e": nil,
	"395bac96-1ce7-4681-99fe-59ed2f6dd156": nil,
	"b79da1ed-4d54-4396-bc8b-38594979cab2": nil,
	"2a5fe0d2-134c-4a4d-bf90-cfafaa795a84": nil,
	"1532f8ae-1dc7-4609-a518-063f8de9c751": nil,
	"350d2dde-2abf-48c0-a1da-63dff2d00bfa": nil,
	"dbf22912-afdd-49cf-ac0a-6417886d8407": nil,
	"36a25a40-d952-4c07-9625-ee88745d6df7": nil,
	"50cf8db2-11e6-44e2-ad50-89b8041ee670": nil,
	"b8303223-d073-42b9-8338-cc3b72e5ae69": nil,
	"dca76df6-bd90-46d2-b032-30560f71de4d": nil,
	"25d36779-1263-4777-aaa1-6c2949addbbf": nil,
	"e0558aef-d2d3-4111-aa79-965cdc604f57": nil,
	"d8248eef-8cc6-498b-83c3-8fd4dd377893": nil,
	"66763699-8e81-458c-a298-20cd843ddd8d": nil,
	"507077aa-7b0b-42c8-a38e-7ff28846c159": nil,
	"fa4fc042-7238-495d-9706-2d7b733bcff7": nil,
	"aadea88f-bf0c-48ac-9ee7-ac6bdfcb819c": nil,
	"f63d041c-e752-4f5c-a3ab-80f45241b249": dm07.KoocPollon,
	"fcb63428-e684-486b-9c34-04d24adc8221": nil,
	"1640acce-f082-4814-be15-b7851f65d21a": nil,
	"190cdfdd-c077-4521-91ed-7bc5d6853b75": nil,
	"fedcde11-394c-4af5-aa48-df0dc9d02647": nil,
	"9f3f83fa-ed3e-48f7-9396-bbe1d79ed544": nil,
	"dd842b81-e187-44e3-92f2-a36826718849": nil,
	"30272397-4b70-4b40-b7ee-8399ebe099d6": nil,
	"e6c53caf-9923-4496-a817-2727ea17bead": nil,
	"8de000dc-4cd9-4a6b-bccc-acfae89211ef": nil,
	"bb3a50dd-049d-488d-b89c-779dcf29b82e": nil,
	"3e3b883e-721c-42b9-8181-52e514695c7f": nil,
	"a5cba5aa-271f-43b0-8350-ca94c5cdfe9a": nil,
	"ced03efd-90c8-4303-83b4-16d2d1198aaf": nil,
}
