package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regexp_match_no_flags(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_match(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_regexp_split_to_table(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int64
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
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == v2 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v22 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v26 = F_pg_detoast_datum_packed(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = v26
					v29 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v31 = int32(_a_F_regexp_split_to_table_0)
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v34
						v37 = v11 + int32(8)
						F_parse_re_flags(m, v37, v28)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
							if v40 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_regexp_split_to_table_1)
										F_errmsg(m, int32(_a_F_regexp_split_to_table_2), v11)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_split_to_table_3), int32(1772), int32(_a_F_regexp_split_to_table_4))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v43 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v43)
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v46 = F_pg_detoast_datum_copy(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int64(0)
								} else {
									v48 = int32(0)
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v51 = int32(1)
									v53 = F_setup_regexp_matches(m, v46, v18, v37, v48, v49, v48, v51, v51)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v32
										*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v53
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
										if v66 <= v67 {
											v69 = F_build_regexp_split_result(m, v65)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int64(0)
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
												v72 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v71 + v72
												v75 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
												*(*int64)(unsafe.Add(mBase, uint32(v64))) = v75 + int64(1)
												v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v72
												v89 = v69
												m.G0 = v11 + int32(16)
												return v89
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int64(0)
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = int32(2)
												v87 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
												v89 = int64(0)
												m.G0 = v11 + int32(16)
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
				v28 = v2
				v29 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = int32(_a_F_regexp_split_to_table_0)
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v34
					v37 = v11 + int32(8)
					F_parse_re_flags(m, v37, v28)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
						if v40 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_regexp_split_to_table_1)
									F_errmsg(m, int32(_a_F_regexp_split_to_table_2), v11)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_split_to_table_3), int32(1772), int32(_a_F_regexp_split_to_table_4))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v43)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v46 = F_pg_detoast_datum_copy(m, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int64(0)
							} else {
								v48 = int32(0)
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v51 = int32(1)
								v53 = F_setup_regexp_matches(m, v46, v18, v37, v48, v49, v48, v51, v51)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v32
									*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v53
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
									if v66 <= v67 {
										v69 = F_build_regexp_split_result(m, v65)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int64(0)
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
											v72 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v71 + v72
											v75 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
											*(*int64)(unsafe.Add(mBase, uint32(v64))) = v75 + int64(1)
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v72
											v89 = v69
											m.G0 = v11 + int32(16)
											return v89
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int64(0)
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = int32(2)
											v87 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
											v89 = int64(0)
											m.G0 = v11 + int32(16)
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
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
		if v66 <= v67 {
			v69 = F_build_regexp_split_result(m, v65)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int64(0)
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
				v72 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v71 + v72
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
				*(*int64)(unsafe.Add(mBase, uint32(v64))) = v75 + int64(1)
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v72
				v89 = v69
				m.G0 = v11 + int32(16)
				return v89
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int64(0)
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v84)+20)) = int32(2)
				v87 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
				v89 = int64(0)
				m.G0 = v11 + int32(16)
				return v89
			}
		}
	}
}
