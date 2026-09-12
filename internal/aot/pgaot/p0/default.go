package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_default_text_search_config(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if base.B2i32(v11 == int32(2)) == int32(0) {
		v93 = int32(1)
		m.G0 = v7 + int32(48)
		return v93
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[226]))
		if v17 == int32(0) {
			v93 = int32(1)
			m.G0 = v7 + int32(48)
			return v93
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[889]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v21
			v24 = *(*int64)(unsafe.Add(mBase, _consts[890]))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v29 = F_stringToQualifiedNameList(m, v26, v7+int32(32))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				if v29 != 0 {
					v34 = F_get_ts_config_oid(m, v29, int32(1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v61 = F_SearchSysCache1(m, int32(74), v34)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								if v61 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v34
										F_errmsg_internal(m, int32(45648), v7+int32(16))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496050), int32(649), int32(335245))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+22)))
									v67 = v65 + v66
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
									v69 = F_get_namespace_name(m, v68)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v73 = F_quote_qualified_identifier(m, v69, v67+int32(4))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v61)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_bms_free(m, v77)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = F_guc_strdup(m, int32(15), v73)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0))) = v81
														F_pfree(m, v73)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															if v86 != 0 {
																v93 = int32(1)
															} else {
																v93 = int32(0)
															}
															m.G0 = v7 + int32(48)
															return v93
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if l2 != int32(12) {
								v93 = base.B2i32(l2 == int32(12))
								m.G0 = v7 + int32(48)
								return v93
							} else {
								v41 = F_errstart(m, int32(18), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									if v41 == int32(0) {
										v93 = base.B2i32(l2 == int32(12))
										m.G0 = v7 + int32(48)
										return v93
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v48
											F_errmsg(m, int32(70995), v7)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496050), int32(635), int32(335245))
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v93 = base.B2i32(l2 == int32(12))
													m.G0 = v7 + int32(48)
													return v93
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l2 != int32(12) {
						v93 = base.B2i32(l2 == int32(12))
						m.G0 = v7 + int32(48)
						return v93
					} else {
						v41 = F_errstart(m, int32(18), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							if v41 == int32(0) {
								v93 = base.B2i32(l2 == int32(12))
								m.G0 = v7 + int32(48)
								return v93
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v48
									F_errmsg(m, int32(70995), v7)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496050), int32(635), int32(335245))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v93 = base.B2i32(l2 == int32(12))
											m.G0 = v7 + int32(48)
											return v93
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_default_multirange_selectivity(m *base.Module, l0 int32) float64 {
	var v5 float64
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v46 float64
	_ = v46
	v5 = float64(0.01)
	if l0 <= int32(4141) {
		v9 = l0 - int32(2862)
		if base.Ui32(int32(15)) < base.Ui32(v9) {
			if l0 == int32(3585) {
				v46 = float64(0.3333333333333333)
			} else {
				if l0 != int32(4035) {
					v46 = v5
				} else {
					v46 = float64(0.3333333333333333)
				}
			}
			return v46
		} else {
			v13 = int32(1) << (uint(v9) % 32)
			if v13&int32(57359) != 0 {
				v46 = float64(0.3333333333333333)
				return v46
			} else {
				if v13&int32(6912) != 0 {
					return float64(0.005)
				} else {
					if int32(1)<<(uint(v9)%32)&int32(1152) == int32(0) {
						if l0 == int32(3585) {
							v46 = float64(0.3333333333333333)
						} else {
							if l0 != int32(4035) {
								v46 = v5
							} else {
								v46 = float64(0.3333333333333333)
							}
						}
						return v46
					} else {
						return float64(0.005)
					}
				}
			}
		}
	} else {
		if base.Ui32(l0-int32(4395)) < base.Ui32(int32(6)) {
			v46 = float64(0.3333333333333333)
			return v46
		} else {
			if base.Ui32(l0-int32(4539)) < base.Ui32(int32(2)) {
				return float64(0.005)
			} else {
				if l0 != int32(4142) {
					v46 = v5
				} else {
					v46 = float64(0.3333333333333333)
				}
				return v46
			}
		}
	}
}
