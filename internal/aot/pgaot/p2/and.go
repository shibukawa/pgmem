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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = F_palloc(m, v10<<(uint(int32(2))%32))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v19 = v17 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
		v22 = l0 + int32(8)
		if v17 != 0 {
			v23 = v22
			v27 = v13
			v28 = v19
			for {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				if v31 != int32(1) {
					v38 = v27
					v39 = v28
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v23
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v38 = v27 + int32(4)
					v39 = v37
				}
				v41 = v39 - int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v41
				if v39 != 0 {
					v23 = v23 + int32(12)
					v27 = v38
					v28 = v41
					continue
				} else {
					break
				}
				break
			}
			v49 = v38
		} else {
			v49 = v13
		}
		v54 = int32(2)
		v55 = (v49 - v13) >> (uint(v54) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v55
		if v54 <= v55 {
			v63 = v22 + v9*int32(12)
			F_qsort_arg(m, v13, v55, int32(4), int32(1535), v63)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if int32(2) <= v66 {
					v71 = v13 + int32(4)
					v74 = v13
					for {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
						v81 = int32(12)
						v84 = int32(4095)
						v85 = v80 & v84
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
						v92 = v87 & v84
						if v85 == int32(0) {
							v98 = int32(0)
							if v98 < v92 {
								v101 = int32(-1)
							} else {
								v101 = v98
							}
							v118 = v101
						} else {
							if v92 == int32(0) {
								v118 = base.B2i32(int32(0) < v85)
							} else {
								if base.Ui32(v85) < base.Ui32(v92) {
									v107 = v85
								} else {
									v107 = v92
								}
								v108 = F_memcmp(m, v63+int32(base.Ui32(v80)>>(uint(v81)%32)), v63+int32(base.Ui32(v87)>>(uint(v81)%32)), v107)
								mBase = m.M
								if v108 != 0 {
									v116 = v108
									v118 = v116
								} else {
									if v85 == v92 {
										v118 = int32(0)
									} else {
										if v85 < v92 {
											v115 = int32(-1)
										} else {
											v115 = int32(1)
										}
										v116 = v115
										v118 = v116
									}
								}
							}
						}
						if v118 != 0 {
							v120 = v74 + int32(4)
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
							*(*int32)(unsafe.Add(mBase, uint32(v120))) = v121
							v123 = v120
						} else {
							v123 = v74
						}
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v126 = v71 + int32(4)
						if (v126-v13)>>(uint(int32(2))%32) < v124 {
							v71 = v126
							v74 = v123
							continue
						} else {
							break
						}
						break
					}
					v134 = v123
				} else {
					v134 = v13
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = (v134 - v13 + int32(4)) >> (uint(int32(2)) % 32)
				return v13
			}
		} else {
			return v13
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v3 = F_text_to_cstring(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_canonicalize_path_enc(m, v3)
		mBase = m.M
		v9 = *(*int32)(unsafe.Add(mBase, _consts[279]))
		v11 = F_has_privs_of_role(m, v9, int32(4569))
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
					v17 = *(*int32)(unsafe.Add(mBase, _consts[203]))
					v19 = F_strlen(m, v17)
					mBase = m.M
					v20 = F_strncmp(m, v17, v3, v19)
					mBase = m.M
					if v20 != 0 {
						v29 = int32(0)
					} else {
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v3))))
						v29 = base.B2i32(v22 == int32(47)) | base.B2i32(v22 == int32(0))
					}
					if v29 != 0 {
						return v3
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
						if v32 == int32(47) {
							v36 = F_strlen(m, v31)
							mBase = m.M
							v37 = F_strncmp(m, v31, v3, v36)
							mBase = m.M
							if v37 != 0 {
								v46 = int32(0)
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v3))))
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
										F_errmsg(m, int32(462141), int32(0))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(525303), int32(84), int32(397720))
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
									F_errmsg(m, int32(462141), int32(0))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(525303), int32(84), int32(397720))
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
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
						if v67 != int32(46) {
							v79 = int32(1)
						} else {
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
							v71 = int32(0)
							v79 = base.B2i32(base.B2i32(v70 == v71)|base.B2i32(v70 == int32(47)) == v71)
						}
					case 1:
						v79 = int32(0)
					default:
						v79 = int32(1)
					}
					if v79&int32(1) != 0 {
						return v3
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(13475), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(525303), int32(89), int32(397720))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
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
