package phenixPhaserInputConfig

var Tmpl = `verbose = False
dry_run = False
show_script = False
show_defaults = False
test_mode = False
phaser {
  mode = ANO CCA EP_AUTO *MR_AUTO MR_FRF MR_FTF MR_PAK MR_RNP NMAXYZ SCEDS
  sad_mode = *SAD MR_SAD
  hklin = "{{.reflectionFile}}"
  labin = I,SIGI
  model = "{{.modelFile}}"
  seq_file = None
  mol_weight = None
  model_rmsd = None
  model_identity = {{.modelIdentity}}
  model_idhi = None
  model_idlo = None
  component_copies = 1
  search_copies = 1
  chain_type = *protein dna rna
  crystal {
    xtal_id = None
    pdb_file = None
    ha_file = None
    dataset {
      wave_id = None
      hklin = None
      labin = None
      wavelength = None
      energy = None
    }
  }
  crystal_symmetry {
    unit_cell = {{.unitcell}}
    space_group = {{.spacegroup}}
  }
  composition {
    solvent = None
    chain {
      chain_type = *protein na
      comp_type = *sequence_file sequence mw nres
      sequence_file = None
      sequence = None
      nres = None
      mw = None
      num = 1
    }
    atom {
      element = None
      num = None
    }
  }
  ensemble {
    model_id = None
    disable_check = False
    use_hetatm = False
    estimator = *oeffner chothialesk
    solution_at_origin = False
    coordinates {
      pdb = None
      rmsd = None
      identity = None
      read_variance_from_pdb_remarks = False
    }
    map_hklin = None
    map_labels = None
    map_extent = None
    map_rms = None
    map_centre = None
    map_protein_mw = 0
    map_nucleic_mw = 0
    map_cell_scale = 1
    ptgroup {
      coverage = 0
      identity = 0
      rmsd = 0
      tolangular = 0
      tolspatial = 0
    }
    bins {
      minimum = 0
      maximum = 0
      width = 0
    }
    trace {
      pdb = None
      sampling {
        use = all calpha hexgrid *auto
        target = None
        distance = None
        range = None
        min = None
        wang = None
        asa = None
        ncycle = None
      }
    }
  }
  search {
    ensembles = None
    copies = 1
  }
  solution = None
  output_dir = "{{.outputDir}}"
  keywords {
    general {
      root = "PHASER"
      title = None
      mute = None
      xyzout = True
      xyzout_ensemble = True
      xyzout_principal = None
      xyzout_packing = False
      xyzout_nma_all = False
      hklout = True
      information = False
      jobs = 4
      keywords = False
    }
    atoms {
      change_scatterer = False
      change_scatterer_type = ""
      change_bfac_wilson = True
      change_original = False
    }
    bfactor_wilson {
      restraint = True
      sigma = 5
    }
    bfactor_sphericity {
      restraint = True
      sigma = 5
    }
    bfactor_refine {
      restraint = True
      sigma = 6
    }
    bins_data {
      minimum = 6
      maximum = 50
      width = 500
    }
    boxscale = 4.0
    composition.minimum_solvent = 20
    composition.percentage = 0.5
    cluster_pdb = None
    eigen {
      read = None
      write = True
    }
    ellg {
      target = 225
      rmsd_mult = None
    }
    formfactors = *xray electron neutron
    find {
      scatterer = "SE"
      cluster = False
      number = 1
      target = "LESSS1"
      special_position = False
      occupancy = 0.25
      purge {
        select = percent *sigma number all
        cutoff = 3.000001
        occupancy = 0.1
      }
      peaks {
        select = *percent sigma number all
        cutoff = 75.000001
        cluster = True
      }
      distance {
        centrosymmetric = -999
        duplicate = -999
      }
    }
    hand = off on *both
    llgcompletion {
      complete = True
      scatterer = None
      clash = 0
      sigma = 6
      holes = True
      top = False
      ncycle = 50
      method = *atomtype imaginary
    }
    llgmaps = False
    macano {
      protocol = *default off custom
      macrocycle {
        ref_anis = None
        ref_bins = None
        ref_solk = None
        ref_solb = None
        ncyc = None
        minimizer = *BFGS NEWTON DESCENT
      }
    }
    macmr {
      protocol = *default off custom
      macrocycle {
        ref_rota = None
        ref_tran = None
        ref_bfac = None
        ref_vrms = None
        ref_cell = None
        ref_ofac = None
        last_only = None
        ncyc = None
        minimizer = *BFGS NEWTON DESCENT
      }
    }
    macsad {
      protocol = *default off custom
      macrocycle {
        ref_pk = None
        ref_pb = None
        ref_xyz = None
        ref_occ = None
        ref_bfac = None
        ref_fdp = None
        ref_sa = None
        ref_sb = None
        ref_sp = None
        ref_sd = None
        ncyc = 50
        target = sadm_only *not_anom_only
        minimizer = *bfgs newton descent
      }
    }
    occupancy {
      min = 0.01
      max = 1
      merge = True
      offset = 0
      fraction = 0.50
      bias = 0.1
      window {
        ellg = 5
        max = 11
        nresidues = 0
      }
    }
    outliers {
      reject = True
      probability = 1e-06
      information = 0.01
    }
    pack {
      select = *percent all
      cutoff = 10
      quick = True
      compact = True
      keep_high_tfz = False
    }
    partial {
      mode = *model map
      pdb = None
      cif = None
      hklin = None
      variance = rms *id
      deviation = None
    }
    partial_labin {
      fc = None
      phic = None
    }
    peaks_rotation {
      select = *percent sigma number all
      cutoff = 75.000001
      cluster = True
      down = 15
    }
    peaks_translation {
      select = *percent sigma number all
      cutoff = 75.000001
      cluster = True
    }
    permutations = False
    purge_rotation {
      enable = True
      percent = 75.000001
      number = 0
    }
    purge_gyre {
      enable = True
      percent = 75.000001
      number = 5
    }
    purge_translation {
      enable = True
      percent = 75.000001
      number = 0
    }
    purge_rnp {
      enable = True
      percent = 75.000001
      number = 0
    }
    rescore_rotation = None
    rescore_translation = None
    resharpen_percent = 100
    resolution {
      high = None
      low = None
      auto_high = None
    }
    rfactor_use = True
    rfactor_cutoff = 40
    rotate {
      volume = *full around
      euler = 0,0,0
      range = 10
    }
    sampling_rotation = None
    sampling_translation = None
    scattering {
      scatterer {
        type = None
        fp = None
        fdp = None
        fixfdp = *EDGE ON OFF
      }
      restraint = True
      sigma = 0.2
    }
    search {
      order_auto = True
      method = *fast full
      pruning = True
      bfactor = 0
      ofactor = 1
    }
    sgalternative {
      select = all *hand list none
      base = None
      test = None
    }
    solparameters {
      siga_fsol = 1.05
      siga_bsol = 501.605
      siga_min = 0.1
      bulk_use = False
      bulk_fsol = 0.35
      bulk_bsol = 45
    }
    solution_origin_ensemble = None
    target_rotation = *fast brute
    target_rotation_type = *LERF1 CORRELATION
    target_translation = *fast brute phased
    target_translation_type = *LETF1 LETF2
    tncs {
      use = True
      nmol = 0
      maxnmol = None
      rlist_add = True
      translation {
        vector = 0,0,0
        perturb = True
      }
      patterson {
        hires = 5.0
        maps = False
        lores = 10.0
        percent = 20
        distance = 15
      }
      variance {
        rmsd = 0.4
        frac = 1.0
      }
      rotation {
        angle = -999,-999,-999
        sampling = -999
        range = -999
      }
      link {
        restraint = True
        sigma = 0.1
      }
      pair_only = False
      gfunction_radius = 0
      commensurate {
        peak = 0.95
        tolf = 0.02
        tolo = 5.0
        patt = 0.1
        miss = 0.2
      }
    }
    topfiles = 1
    translation {
      volume = *full line region around
      fractional = True
      start = 0,0,0
      end = 1,1,1
      range = 5
      packing_use = True
      packing_base = None
      packing_cutoff = 50
      packing_number = 500
    }
    output_level = *logfile silent concise summary verbose debug
    zscore {
      use = True
      stop = False
      solved = 8
      possibly_solved = 6
      half = True
    }
    perturb {
      increment = None
      rms {
        step = None
        maximum = None
        stretch = None
        clash = None
        direction = *tofro backwards forwards
      }
    }
    nma {
      mode = None
      combinations = None
      original = None
    }
  }
  run_control {
    gui_run = False
  }
}`
