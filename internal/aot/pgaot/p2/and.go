package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SortAndUniqItems(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = F_palloc_mul(m, int32(4), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v18 = v16 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v18
		v21 = l0 + int32(8)
		if v16 != 0 {
			v22 = v21
			v24 = v12
			v27 = v18
			for {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				if v30 != int32(1) {
					v37 = v24
					v38 = v27
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = v22
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v37 = v24 + int32(4)
					v38 = v36
				}
				v40 = v38 - int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
				if v38 != 0 {
					v22 = v22 + int32(12)
					v24 = v37
					v27 = v40
					continue
				} else {
					break
				}
				break
			}
			v46 = v37
		} else {
			v46 = v12
		}
		v53 = int32(2)
		v54 = (v46 - v12) >> (uint(v53) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v54
		if v53 <= v54 {
			v62 = v21 + v9*int32(12)
			F_qsort_arg(m, v12, v54, int32(4), int32(1729), v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if int32(2) <= v65 {
					v70 = v12 + int32(4)
					v72 = v12
					for {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
						v80 = int32(12)
						v83 = int32(4095)
						v84 = v79 & v83
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
						v91 = v86 & v83
						if v84 == int32(0) {
							v97 = int32(0)
							if v97 < v91 {
								v100 = int32(-1)
							} else {
								v100 = v97
							}
							v117 = v100
						} else {
							if v91 == int32(0) {
								v117 = base.B2i32(int32(0) < v84)
							} else {
								if base.Ui32(v84) < base.Ui32(v91) {
									v106 = v84
								} else {
									v106 = v91
								}
								v107 = F_memcmp(m, v62+int32(base.Ui32(v79)>>(uint(v80)%32)), v62+int32(base.Ui32(v86)>>(uint(v80)%32)), v106)
								mBase = m.M
								if v107 != 0 {
									v115 = v107
									v117 = v115
								} else {
									if v84 == v91 {
										v117 = int32(0)
									} else {
										if v84 < v91 {
											v114 = int32(-1)
										} else {
											v114 = int32(1)
										}
										v115 = v114
										v117 = v115
									}
								}
							}
						}
						if v117 != 0 {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v118
							v122 = v72 + int32(4)
						} else {
							v122 = v72
						}
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v125 = v70 + int32(4)
						if (v125-v12)>>(uint(int32(2))%32) < v123 {
							v70 = v125
							v72 = v122
							continue
						} else {
							break
						}
						break
					}
					v132 = v122
				} else {
					v132 = v12
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = (v132 - v12 + int32(4)) >> (uint(int32(2)) % 32)
				return v12
			}
		} else {
			return v12
		}
	}
}
func F_convert_and_check_filename(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v3 = F_text_to_cstring(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_canonicalize_path_enc(m, v3)
		mBase = m.M
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_convert_and_check_filename[0]))
		v11 = F_has_privs_of_role(m, v9, int32(_a_F_convert_and_check_filename_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v11 != 0 {
				return v3
			} else {
				v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
				if v13 == int32(47) {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_convert_and_check_filename[1]))
					v18 = F_strlen(m, v17)
					mBase = m.M
					v19 = F_strncmp(m, v17, v3, v18)
					mBase = m.M
					if v19 != 0 {
						v29 = int32(0)
					} else {
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v3))))
						v29 = base.B2i32(v22 == int32(47)) | base.B2i32(v22 == int32(0))
					}
					if v29 != 0 {
						return v3
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_convert_and_check_filename[2]))
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
						if v32 == int32(47) {
							v35 = F_strlen(m, v31)
							mBase = m.M
							v36 = F_strncmp(m, v31, v3, v35)
							mBase = m.M
							if v36 != 0 {
								v46 = int32(0)
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v3))))
								v46 = base.B2i32(v39 == int32(47)) | base.B2i32(v39 == int32(0))
							}
							if v46 != 0 {
								return v3
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_convert_and_check_filename_1), int32(0))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_convert_and_check_filename_2), int32(84), int32(_a_F_convert_and_check_filename_3))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_convert_and_check_filename_1), int32(0))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_convert_and_check_filename_2), int32(84), int32(_a_F_convert_and_check_filename_3))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				} else {
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
					switch v63 - int32(46) {
					case 0:
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
						if v66 != int32(46) {
							v78 = int32(1)
						} else {
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
							v70 = int32(0)
							v78 = base.B2i32(base.B2i32(v69 == v70)|base.B2i32(v69 == int32(47)) == v70)
						}
					case 1:
						v78 = int32(0)
					default:
						v78 = int32(1)
					}
					if v78 != 0 {
						return v3
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_convert_and_check_filename_4), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_convert_and_check_filename_2), int32(89), int32(_a_F_convert_and_check_filename_3))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
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
