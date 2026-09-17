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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == v2 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v21 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v25 = F_pg_detoast_datum_packed(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					v28 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = int32(_a_F_regexp_split_to_table_0)
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0]))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v33
						v36 = v10 + int32(8)
						F_parse_re_flags(m, v36, v27)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
							if v39 == int32(1) {
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
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_regexp_split_to_table_1)
										F_errmsg(m, int32(_a_F_regexp_split_to_table_2), v10)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_regexp_split_to_table_3), int32(1771), int32(_a_F_regexp_split_to_table_4))
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
								v42 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v42)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v45 = F_pg_detoast_datum_copy(m, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v47 = int32(0)
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v50 = int32(1)
									v52 = F_setup_regexp_matches(m, v45, v17, v36, v47, v48, v47, v50, v50)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v31
										*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v52
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
												m.G0 = v10 + int32(16)
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
												m.G0 = v10 + int32(16)
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
				v27 = v2
				v28 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(_a_F_regexp_split_to_table_0)
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0]))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v33
					v36 = v10 + int32(8)
					F_parse_re_flags(m, v36, v27)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						if v39 == int32(1) {
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
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_regexp_split_to_table_1)
									F_errmsg(m, int32(_a_F_regexp_split_to_table_2), v10)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_split_to_table_3), int32(1771), int32(_a_F_regexp_split_to_table_4))
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
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v42)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v45 = F_pg_detoast_datum_copy(m, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = int32(0)
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v50 = int32(1)
								v52 = F_setup_regexp_matches(m, v45, v17, v36, v47, v48, v47, v50, v50)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_regexp_split_to_table[0])) = v31
									*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v52
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
											m.G0 = v10 + int32(16)
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
											m.G0 = v10 + int32(16)
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
				m.G0 = v10 + int32(16)
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
				m.G0 = v10 + int32(16)
				return v89
			}
		}
	}
}
