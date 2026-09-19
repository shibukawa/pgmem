package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UpdateChangedParamSet(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+68))
	v5 = F_bms_intersect(m, v4, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v8 = F_bms_join(m, v7, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v8
			return
		}
	}
}
func F_unlink_span(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v42 = v10
		if v42 != 0 {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+1468))
			if v46 != v48 {
				v53 = F_LWLockAcquire(m, v47+int32(1476), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v57+int32(1476))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
							v68 = l0 + v65*int32(20)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
							if v69 != 0 {
								v73 = v69
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
								return
							} else {
								v70 = F_get_segment_by_index(m, l0, v65)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
									v73 = v72
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
									return
								}
							}
						}
					}
				}
			} else {
				v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
				v68 = l0 + v65*int32(20)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
				if v69 != 0 {
					v73 = v69
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
					return
				} else {
					v70 = F_get_segment_by_index(m, l0, v65)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
						v73 = v72
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
						return
					}
				}
			}
		} else {
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+1468))
			if v78 != v80 {
				v85 = F_LWLockAcquire(m, v79+int32(1476), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v89+int32(1476))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
							v100 = l0 + v97*int32(20)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
							if v101 != 0 {
								v105 = v101
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
								return
							} else {
								v102 = F_get_segment_by_index(m, l0, v97)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
									v105 = v104
									v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
									v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
									return
								}
							}
						}
					}
				}
			} else {
				v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
				v100 = l0 + v97*int32(20)
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
				if v101 != 0 {
					v105 = v101
					v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
					return
				} else {
					v102 = F_get_segment_by_index(m, l0, v97)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
						v105 = v104
						v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
						return
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1468))
		if v11 != v13 {
			v18 = F_LWLockAcquire(m, v12+int32(1476), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v22+int32(1476))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v30 = int32(base.Ui32(v7) >> (uint(int32(27)) % 32))
						v33 = l0 + v30*int32(20)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						if v34 != 0 {
							v38 = v34
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v38+v7&int32(134217727))+4)) = v40
							v42 = v40
							if v42 != 0 {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+1468))
								if v46 != v48 {
									v53 = F_LWLockAcquire(m, v47+int32(1476), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v57+int32(1476))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
												v68 = l0 + v65*int32(20)
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
												if v69 != 0 {
													v73 = v69
													v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
													return
												} else {
													v70 = F_get_segment_by_index(m, l0, v65)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
														v73 = v72
														v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
														return
													}
												}
											}
										}
									}
								} else {
									v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
									v68 = l0 + v65*int32(20)
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
									if v69 != 0 {
										v73 = v69
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
										return
									} else {
										v70 = F_get_segment_by_index(m, l0, v65)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
											v73 = v72
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
											return
										}
									}
								}
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+1468))
								if v78 != v80 {
									v85 = F_LWLockAcquire(m, v79+int32(1476), int32(0))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v89+int32(1476))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
												v100 = l0 + v97*int32(20)
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
												if v101 != 0 {
													v105 = v101
													v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
													v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
													return
												} else {
													v102 = F_get_segment_by_index(m, l0, v97)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
														v105 = v104
														v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
														v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
														return
													}
												}
											}
										}
									}
								} else {
									v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
									v100 = l0 + v97*int32(20)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
									if v101 != 0 {
										v105 = v101
										v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
										return
									} else {
										v102 = F_get_segment_by_index(m, l0, v97)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
											v105 = v104
											v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
											return
										}
									}
								}
							}
						} else {
							v35 = F_get_segment_by_index(m, l0, v30)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
								v38 = v37
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v38+v7&int32(134217727))+4)) = v40
								v42 = v40
								if v42 != 0 {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+1468))
									if v46 != v48 {
										v53 = F_LWLockAcquire(m, v47+int32(1476), int32(0))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v57+int32(1476))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
													v68 = l0 + v65*int32(20)
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
													if v69 != 0 {
														v73 = v69
														v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
														return
													} else {
														v70 = F_get_segment_by_index(m, l0, v65)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
															v73 = v72
															v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
															return
														}
													}
												}
											}
										}
									} else {
										v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
										v68 = l0 + v65*int32(20)
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
										if v69 != 0 {
											v73 = v69
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
											return
										} else {
											v70 = F_get_segment_by_index(m, l0, v65)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
												v73 = v72
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
												return
											}
										}
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+1468))
									if v78 != v80 {
										v85 = F_LWLockAcquire(m, v79+int32(1476), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v89+int32(1476))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
													v100 = l0 + v97*int32(20)
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
													if v101 != 0 {
														v105 = v101
														v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
														v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
														return
													} else {
														v102 = F_get_segment_by_index(m, l0, v97)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
															v105 = v104
															v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
															v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
															return
														}
													}
												}
											}
										}
									} else {
										v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
										v100 = l0 + v97*int32(20)
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
										if v101 != 0 {
											v105 = v101
											v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
											return
										} else {
											v102 = F_get_segment_by_index(m, l0, v97)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return
											} else {
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
												v105 = v104
												v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
												v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
												return
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
			v30 = int32(base.Ui32(v7) >> (uint(int32(27)) % 32))
			v33 = l0 + v30*int32(20)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
			if v34 != 0 {
				v38 = v34
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v38+v7&int32(134217727))+4)) = v40
				v42 = v40
				if v42 != 0 {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+1468))
					if v46 != v48 {
						v53 = F_LWLockAcquire(m, v47+int32(1476), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v57+int32(1476))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
									v68 = l0 + v65*int32(20)
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
									if v69 != 0 {
										v73 = v69
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
										return
									} else {
										v70 = F_get_segment_by_index(m, l0, v65)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
											v73 = v72
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
											return
										}
									}
								}
							}
						}
					} else {
						v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
						v68 = l0 + v65*int32(20)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
						if v69 != 0 {
							v73 = v69
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
							return
						} else {
							v70 = F_get_segment_by_index(m, l0, v65)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
								v73 = v72
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
								return
							}
						}
					}
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+1468))
					if v78 != v80 {
						v85 = F_LWLockAcquire(m, v79+int32(1476), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v89+int32(1476))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
									v100 = l0 + v97*int32(20)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
									if v101 != 0 {
										v105 = v101
										v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
										return
									} else {
										v102 = F_get_segment_by_index(m, l0, v97)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
											v105 = v104
											v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
											return
										}
									}
								}
							}
						}
					} else {
						v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
						v100 = l0 + v97*int32(20)
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
						if v101 != 0 {
							v105 = v101
							v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
							return
						} else {
							v102 = F_get_segment_by_index(m, l0, v97)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
								v105 = v104
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
								return
							}
						}
					}
				}
			} else {
				v35 = F_get_segment_by_index(m, l0, v30)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					v38 = v37
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v38+v7&int32(134217727))+4)) = v40
					v42 = v40
					if v42 != 0 {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+1468))
						if v46 != v48 {
							v53 = F_LWLockAcquire(m, v47+int32(1476), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v57+int32(1476))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
										v68 = l0 + v65*int32(20)
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
										if v69 != 0 {
											v73 = v69
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
											return
										} else {
											v70 = F_get_segment_by_index(m, l0, v65)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
												v73 = v72
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
												return
											}
										}
									}
								}
							}
						} else {
							v65 = int32(base.Ui32(v42) >> (uint(int32(27)) % 32))
							v68 = l0 + v65*int32(20)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
							if v69 != 0 {
								v73 = v69
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
								return
							} else {
								v70 = F_get_segment_by_index(m, l0, v65)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
									v73 = v72
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v73+v42&int32(134217727))+8)) = v75
									return
								}
							}
						}
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+1468))
						if v78 != v80 {
							v85 = F_LWLockAcquire(m, v79+int32(1476), int32(0))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v89+int32(1476))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
										v100 = l0 + v97*int32(20)
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
										if v101 != 0 {
											v105 = v101
											v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
											return
										} else {
											v102 = F_get_segment_by_index(m, l0, v97)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return
											} else {
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
												v105 = v104
												v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
												v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
												return
											}
										}
									}
								}
							}
						} else {
							v97 = int32(base.Ui32(v77) >> (uint(int32(27)) % 32))
							v100 = l0 + v97*int32(20)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
							if v101 != 0 {
								v105 = v101
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
								return
							} else {
								v102 = F_get_segment_by_index(m, l0, v97)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
									v105 = v104
									v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
									v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v105+v77&int32(134217727)+v107<<(uint(int32(2))%32))+16)) = v111
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_upc_cast_from_ean13(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13948(m, l0, int32(6))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_upc_in(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13891(m, l0, int32(6))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_updateAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	var v11 int32
	_ = v11
	F_updateAclDependenciesWorker(m, l0, l1, l2, l3, int32(97), l4, l5, l6, l7)
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_updateAclDependenciesWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v382 int32
	_ = v382
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v438 int32
	_ = v438
	var v451 int32
	_ = v451
	var v468 int32
	_ = v468
	var v497 int32
	_ = v497
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v560 int32
	_ = v560
	var v572 int32
	_ = v572
	var v585 int32
	_ = v585
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v620 int32
	_ = v620
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	v10 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	if l5 <= v10 {
		v213 = v10
		v214 = v10
		v216 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v213 < l7 {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	v34 = v10
	v35 = v10
	v36 = v10
	v38 = v10
	goto L3
L3:
	;
	if v35 < l7 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if l5 <= v79 {
		v213 = v80
		v214 = v81
		v216 = v83
		goto L1
	} else {
		goto L16
	}
L5:
	;
	v44 = int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l6+v34<<(uint(v44)%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l8+v35<<(uint(v44)%32))))
	if v47 == v51 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v79 = v34
	v80 = v35
	v81 = v36
	v83 = v38
	goto L7
L7:
	;
	goto L4
L8:
	;
	if v74 < l5 {
		v34 = v74
		v35 = v75
		v36 = v76
		v38 = v77
		goto L3
	} else {
		goto L15
	}
L9:
	;
	v53 = int32(1)
	v74 = v34 + v53
	v75 = v35 + v53
	v76 = v36
	v77 = v38
	goto L8
L10:
	;
	goto L11
L11:
	;
	if base.Ui32(v47) < base.Ui32(v51) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v38<<(uint(int32(2))%32)))) = v47
	v62 = int32(1)
	v74 = v34 + v62
	v75 = v35
	v76 = v36
	v77 = v38 + v62
	goto L8
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8+v36<<(uint(int32(2))%32)))) = v51
	v70 = int32(1)
	v74 = v34
	v75 = v35 + v70
	v76 = v36 + v70
	v77 = v38
	goto L8
L15:
	;
	v79 = v74
	v80 = v75
	v81 = v76
	v83 = v77
	goto L7
L16:
	;
	v89 = (l5 - v79) & int32(3)
	if v89 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v144 = l5 + v83 - v79
	if base.Ui32(v79-l5) < base.Ui32(int32(-3)) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v138 = v79
	v139 = v83
	goto L17
L19:
	;
	goto L20
L20:
	;
	v105 = v79
	v106 = v83
	v108 = int32(0)
	goto L21
L21:
	;
	v111 = int32(2)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l6+v105<<(uint(v111)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l6+v106<<(uint(v111)%32)))) = v117
	v119 = int32(1)
	v120 = v106 + v119
	v122 = v105 + v119
	v124 = v108 + v119
	if v124 != v89 {
		v105 = v122
		v106 = v120
		v108 = v124
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v138 = v122
	v139 = v120
	goto L17
L23:
	;
	goto L22
L24:
	;
	v160 = v138
	v161 = v139
	goto L27
L25:
	;
	goto L26
L26:
	;
	v213 = v80
	v214 = v81
	v216 = v144
	goto L1
L27:
	;
	v166 = int32(2)
	v168 = l6 + v161<<(uint(v166)%32)
	v171 = l6 + v160<<(uint(v166)%32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v174
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v178
	v180 = int32(4)
	v183 = v161 + v180
	if v183 != v144 {
		v160 = v160 + v180
		v161 = v183
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	goto L28
L30:
	;
	v225 = (l7 - v213) & int32(3)
	if v225 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v350 = v214
	goto L32
L32:
	;
	v357 = int32(0)
	if base.B2i32(v216 <= v357)&base.B2i32(v350 <= v357) == v357 {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v280 = l7 + v214 - v213
	if base.Ui32(v213-l7) <= base.Ui32(int32(-4)) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v271 = v213
	v273 = v214
	goto L33
L35:
	;
	goto L36
L36:
	;
	v238 = v213
	v240 = v214
	v241 = int32(0)
	goto L37
L37:
	;
	v247 = int32(2)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l8+v238<<(uint(v247)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l8+v240<<(uint(v247)%32)))) = v253
	v255 = int32(1)
	v256 = v240 + v255
	v258 = v238 + v255
	v260 = v241 + v255
	if v260 != v225 {
		v238 = v258
		v240 = v256
		v241 = v260
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v271 = v258
	v273 = v256
	goto L33
L39:
	;
	goto L38
L40:
	;
	v293 = v271
	v295 = v273
	goto L43
L41:
	;
	goto L42
L42:
	;
	v350 = v280
	goto L32
L43:
	;
	v302 = int32(2)
	v304 = l8 + v295<<(uint(v302)%32)
	v307 = l8 + v293<<(uint(v302)%32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+4)) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+8)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = v314
	v316 = int32(4)
	v319 = v295 + v316
	if v319 != v280 {
		v293 = v293 + v316
		v295 = v319
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v366 = F_table_open(m, int32(1214), int32(3))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if l6 != 0 {
		goto L105
	} else {
		goto L106
	}
L49:
	;
	return
L50:
	;
	if int32(0) < v350 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v382 = int32(0)
	goto L54
L52:
	;
	goto L53
L53:
	;
	if int32(0) < v216 {
		goto L93
	} else {
		goto L94
	}
L54:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l8+v382<<(uint(int32(2))%32))))
	if base.B2i32(base.B2i32(l4 != int32(97)) == int32(0))&base.B2i32(v396 == l3) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v524 = v382 + int32(1)
	if v524 != v350 {
		v382 = v524
		goto L54
	} else {
		goto L92
	}
L57:
	;
	v409 = int32(1)
	goto L58
L58:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_updateAclDependenciesWorker_0)) < base.Ui32(v396)) == int32(0))&((v409|base.B2i32(v396 != int32(2200)))&v409) != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	F_shdepLockAndCheckObject(m, int32(1260), v396)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v420 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+11)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v420
	v426 = int32(1)
	if l0 <= int32(3591) {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(1260)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = l0
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_updateAclDependenciesWorker[0]))
	if v497 != 0 {
		goto L86
	} else {
		goto L87
	}
L62:
	;
	goto L61
L63:
	;
	v497 = int32(0)
	goto L62
L64:
	;
	if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v497 = v426
		goto L62
	} else {
		goto L85
	}
L65:
	;
	if l0 <= int32(2670) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if l0 <= int32(_a_F_updateAclDependenciesWorker_1) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	switch l0 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v497 = v426
		goto L62
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L63
	default:
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v438 = l0 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v438))|base.B2i32(int32(1)<<(uint(v438)%32)&int32(226492515) == int32(0)) != 0 {
		goto L64
	} else {
		goto L73
	}
L71:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v497 = v426
	goto L62
L73:
	;
	v497 = v426
	goto L62
L74:
	;
	if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
		v497 = v426
		goto L62
	} else {
		goto L83
	}
L75:
	;
	v451 = l0 - int32(_a_F_updateAclDependenciesWorker_2)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v451))|base.B2i32(int32(1)<<(uint(v451)%32)&int32(963) == int32(0)) != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	switch l0 - int32(_a_F_updateAclDependenciesWorker_3) {
	case 0, 1, 2, 3, 4, 59, 60:
		v497 = v426
		goto L62
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L63
	default:
		goto L79
	}
L78:
	;
	v497 = v426
	goto L62
L79:
	;
	if base.Ui32(l0-int32(_a_F_updateAclDependenciesWorker_4)) < base.Ui32(int32(3)) {
		v497 = v426
		goto L62
	} else {
		goto L80
	}
L80:
	;
	v468 = l0 - int32(_a_F_updateAclDependenciesWorker_5)
	if base.Ui32(int32(15)) < base.Ui32(v468) {
		goto L63
	} else {
		goto L81
	}
L81:
	;
	if int32(1)<<(uint(v468)%32)&int32(_a_F_updateAclDependenciesWorker_6) != 0 {
		v497 = v426
		goto L62
	} else {
		goto L82
	}
L82:
	;
	goto L63
L83:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
		goto L63
	} else {
		goto L84
	}
L84:
	;
	v497 = v426
	goto L62
L85:
	;
	goto L63
L86:
	;
	v508 = int32(0)
	goto L88
L87:
	;
	v508 = v507
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v366)+52))
	v515 = F_heap_form_tuple(m, v510, v21+int32(16), v21+int32(8))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L49
	} else {
		goto L89
	}
L89:
	;
	F_CatalogTupleInsert(m, v366, v515)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L49
	} else {
		goto L90
	}
L90:
	;
	F_pfree(m, v515)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L49
	} else {
		goto L91
	}
L91:
	;
	goto L56
L92:
	;
	goto L55
L93:
	;
	v560 = int32(0)
	goto L96
L94:
	;
	goto L95
L95:
	;
	F_relation_close(m, v366, int32(3))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L49
	} else {
		goto L104
	}
L96:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l6+v560<<(uint(int32(2))%32))))
	if base.B2i32(base.B2i32(l4 != int32(97)) == int32(0))&base.B2i32(v572 == l3) != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L95
L98:
	;
	v598 = v560 + int32(1)
	if v598 != v216 {
		v560 = v598
		goto L96
	} else {
		goto L103
	}
L99:
	;
	v585 = int32(1)
	goto L100
L100:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_updateAclDependenciesWorker_0)) < base.Ui32(v572)) == int32(0))&((v585|base.B2i32(v572 != int32(2200)))&v585) != 0 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	F_shdepDropDependency(m, v366, l0, l1, l2, int32(0), int32(1260), v572, l4)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L49
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	goto L97
L104:
	;
	goto L48
L105:
	;
	F_pfree(m, l6)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L49
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if l8 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	F_pfree(m, l8)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L49
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	m.G0 = v21 + int32(48)
	return
L112:
	;
	goto L111
}
func F_update_metainfo_datafile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v9&int32(25) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(144)
	return
L2:
	;
	v15 = F_unlink(m, int32(_a_F_update_metainfo_datafile_0))
	mBase = m.M
	if int32(0) <= v15 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[2]))
	v42 = F_umask(m, v41)
	mBase = m.M
	v45 = F_fopen(m, int32(_a_F_update_metainfo_datafile_4), int32(_a_F_update_metainfo_datafile_5))
	mBase = m.M
	v46 = F_umask(m, v42)
	mBase = m.M
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[1]))
	if v19 == int32(44) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v24 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_update_metainfo_datafile_0)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_1), v6)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1490), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L1
L13:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[4]))
	if v115 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+80)) = int32(-1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v96 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L33
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[3]))
	if v57 == int32(0) {
		goto L13
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+80)) = int32(10)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v52 | int32(64)
	goto L17
L21:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v61&int32(1) == int32(0) {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+128)) = v57
	v70 = F_pg_fprintf(m, v45, int32(_a_F_update_metainfo_datafile_10), v6+int32(128))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	if int32(0) <= v70 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v76 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v92 = F_fclose(m, v45)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L32
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_8), v6+int32(112))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1524), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	goto L1
L33:
	;
	if v96 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_11), v6+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1513), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L1
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[5]))
	if v153 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v119&int32(8) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v115
	v128 = F_pg_fprintf(m, v45, int32(_a_F_update_metainfo_datafile_9), v6+int32(96))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if int32(0) <= v128 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v134 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	if v134 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v150 = F_fclose(m, v45)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L50
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_8), v6+int32(80))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1537), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L1
L51:
	;
	v190 = F_fclose(m, v45)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L64
	}
L52:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_update_metainfo_datafile[0])))
	if v157&int32(16) == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v153
	v166 = F_pg_fprintf(m, v45, int32(_a_F_update_metainfo_datafile_7), v6-int32(-64))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	if int32(0) <= v166 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v172 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	if v172 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v188 = F_fclose(m, v45)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L63
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_8), v6+int32(48))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1550), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L1
L64:
	;
	v194 = F_rename(m, int32(_a_F_update_metainfo_datafile_4), int32(_a_F_update_metainfo_datafile_0))
	mBase = m.M
	if v194 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v199 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	if v199 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(_a_F_update_metainfo_datafile_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_update_metainfo_datafile_4)
	F_errmsg(m, int32(_a_F_update_metainfo_datafile_6), v6+int32(32))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_update_metainfo_datafile_2), int32(1561), int32(_a_F_update_metainfo_datafile_3))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	goto L1
}
func F_utime(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v10 = m.Env.X__syscall_utimensat(m, int32(-100), l0, v2, v2)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v10) {
		*(*int32)(unsafe.Add(mBase, _c_F_utime[0])) = int32(0) - v10
	} else {
	}
	m.G0 = v5 + int32(32)
	return
}
