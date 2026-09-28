package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetDefaultCharSignedness(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetDefaultCharSignedness[0]))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+272)))
	return v3
}
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_text_search_config[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if base.B2i32(v11 == int32(2)) == int32(0) {
		v94 = int32(1)
		m.G0 = v7 + int32(48)
		return v94
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_text_search_config[1]))
		if v17 == int32(0) {
			v94 = int32(1)
			m.G0 = v7 + int32(48)
			return v94
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_text_search_config[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v21
			v24 = *(*int64)(unsafe.Add(mBase, _c_F_check_default_text_search_config[3]))
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
							v62 = F_SearchSysCache1(m, int32(74), base.I64_extend_i32_u(v34))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								if v62 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v34
										F_errmsg_internal(m, int32(_a_F_check_default_text_search_config_0), v7+int32(16))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_check_default_text_search_config_1), int32(665), int32(_a_F_check_default_text_search_config_2))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
									v68 = v66 + v67
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
									v70 = F_get_namespace_name(m, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v74 = F_quote_qualified_identifier(m, v70, v68+int32(4))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v62)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_bms_free(m, v78)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v82 = F_guc_strdup(m, int32(15), v74)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0))) = v82
														F_pfree(m, v74)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															if v87 != 0 {
																v94 = int32(1)
															} else {
																v94 = int32(0)
															}
															m.G0 = v7 + int32(48)
															return v94
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
								v94 = base.B2i32(l2 == int32(12))
								m.G0 = v7 + int32(48)
								return v94
							} else {
								v41 = F_errstart(m, int32(18), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									if v41 == int32(0) {
										v94 = base.B2i32(l2 == int32(12))
										m.G0 = v7 + int32(48)
										return v94
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v48
											F_errmsg(m, int32(_a_F_check_default_text_search_config_3), v7)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_check_default_text_search_config_1), int32(651), int32(_a_F_check_default_text_search_config_2))
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v94 = base.B2i32(l2 == int32(12))
													m.G0 = v7 + int32(48)
													return v94
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
						v94 = base.B2i32(l2 == int32(12))
						m.G0 = v7 + int32(48)
						return v94
					} else {
						v41 = F_errstart(m, int32(18), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							if v41 == int32(0) {
								v94 = base.B2i32(l2 == int32(12))
								m.G0 = v7 + int32(48)
								return v94
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v48
									F_errmsg(m, int32(_a_F_check_default_text_search_config_3), v7)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_check_default_text_search_config_1), int32(651), int32(_a_F_check_default_text_search_config_2))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v94 = base.B2i32(l2 == int32(12))
											m.G0 = v7 + int32(48)
											return v94
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
	var v4 float64
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v41 float64
	_ = v41
	v4 = float64(0.005)
	if l0 <= int32(_a_F_default_multirange_selectivity_0) {
		v8 = l0 - int32(2862)
		if base.Ui32(int32(15)) < base.Ui32(v8) {
			if base.B2i32(l0 == int32(3585))|base.B2i32(l0 == int32(4035)) != 0 {
				v41 = float64(0.3333333333333333)
				return v41
			} else {
				return float64(0.01)
			}
		} else {
			v12 = int32(1) << (uint(v8) % 32)
			if v12&int32(_a_F_default_multirange_selectivity_1) != 0 {
				v41 = float64(0.3333333333333333)
				return v41
			} else {
				if v12&int32(_a_F_default_multirange_selectivity_2) == int32(0) {
					if base.B2i32(l0 == int32(3585))|base.B2i32(l0 == int32(4035)) != 0 {
						v41 = float64(0.3333333333333333)
						return v41
					} else {
						return float64(0.01)
					}
				} else {
					v41 = v4
					return v41
				}
			}
		}
	} else {
		if base.Ui32(l0-int32(_a_F_default_multirange_selectivity_3)) < base.Ui32(int32(6)) {
			v41 = float64(0.3333333333333333)
			return v41
		} else {
			if base.Ui32(l0-int32(_a_F_default_multirange_selectivity_4)) < base.Ui32(int32(2)) {
				v41 = v4
				return v41
			} else {
				if l0 != int32(_a_F_default_multirange_selectivity_5) {
					return float64(0.01)
				} else {
					v41 = float64(0.3333333333333333)
					return v41
				}
			}
		}
	}
}
