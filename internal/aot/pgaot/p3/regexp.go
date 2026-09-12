package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_match_no_flags(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_match(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_regexp_split_to_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == v2 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v20 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v24
					v27 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(4442576)
						v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v32
						F_parse_re_flags(m, v9+int32(8), v26)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
							if v38 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(642504)
										F_errmsg(m, int32(234187), v9)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(472485), int32(1771), int32(373081))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v41)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v44 = F_pg_detoast_datum_copy(m, v43)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v48 = int32(0)
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v51 = int32(1)
									v53 = F_setup_regexp_matches(m, v44, v16, v9+int32(8), v48, v49, v48, v51, v51)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
										*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v53
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										if v65 <= v66 {
											v68 = F_build_regexp_split_result(m, v64)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
												v71 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v70 + v71
												v74 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
												*(*int64)(unsafe.Add(mBase, uint32(v63))) = v74 + int64(1)
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v71
												v89 = v68
												m.G0 = v9 + int32(16)
												return v89
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(2)
												v86 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v86)
												v89 = int32(0)
												m.G0 = v9 + int32(16)
												return v89
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = v2
				v27 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(4442576)
					v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v32
					F_parse_re_flags(m, v9+int32(8), v26)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
						if v38 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(642504)
									F_errmsg(m, int32(234187), v9)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472485), int32(1771), int32(373081))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v41 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v41)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v44 = F_pg_detoast_datum_copy(m, v43)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v48 = int32(0)
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v51 = int32(1)
								v53 = F_setup_regexp_matches(m, v44, v16, v9+int32(8), v48, v49, v48, v51, v51)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
									*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v53
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
									if v65 <= v66 {
										v68 = F_build_regexp_split_result(m, v64)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
											v71 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v70 + v71
											v74 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
											*(*int64)(unsafe.Add(mBase, uint32(v63))) = v74 + int64(1)
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v71
											v89 = v68
											m.G0 = v9 + int32(16)
											return v89
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(2)
											v86 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v86)
											v89 = int32(0)
											m.G0 = v9 + int32(16)
											return v89
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
		if v65 <= v66 {
			v68 = F_build_regexp_split_result(m, v64)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
				v71 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v70 + v71
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
				*(*int64)(unsafe.Add(mBase, uint32(v63))) = v74 + int64(1)
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v71
				v89 = v68
				m.G0 = v9 + int32(16)
				return v89
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(2)
				v86 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v86)
				v89 = int32(0)
				m.G0 = v9 + int32(16)
				return v89
			}
		}
	}
}
